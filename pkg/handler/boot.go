package handler

import (
	"fmt"
	"net/http"

	"github.com/futuretea/harvester-netboot/pkg/config"
	"github.com/futuretea/harvester-netboot/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const (
	BootPath = "/v1/boot/:mac"
)

func BootHandler(c *gin.Context) {
	mac := c.Param("mac")
	log.Debug().Str("mac", mac).Msg("bootHandler")

	resp, err := genResp(mac)
	if err != nil {
		log.Error().Err(err).Msg("bootHandler")
		return
	}

	log.Debug().Interface("resp", resp).Msg("bootHandler")
	c.JSON(http.StatusOK, resp)
}

func genResp(mac string) (gin.H, error) {
	for _, cluster := range config.Conf.Clusters {
		for _, node := range cluster.Nodes {
			if node.Mac == mac {
				return fmtResp(cluster, node)
			}
		}
	}
	return gin.H{}, nil
}

func fmtConfigFileName(mac string) string {
	return fmt.Sprintf("/tmp/harvester-netboot-generated-%s-config.yaml", mac)
}

func fmtConfigURL(configFileName string) string {
	return fmt.Sprintf("{{ URL \"%s\" }}", configFileName)
}

func fmtKernel(baseURL string, version string) string {
	return fmt.Sprintf("%s/%s/harvester-%s-vmlinuz-amd64", baseURL, version, version)
}

func fmtInitrd(baseURL string, version string) string {
	return fmt.Sprintf("%s/%s/harvester-%s-initrd-amd64", baseURL, version, version)
}

func fmtRootLive(baseURL string, version string) string {
	return fmt.Sprintf("%s/%s/harvester-%s-rootfs-amd64.squashfs", baseURL, version, version)
}

func fmtCmdline(baseURL string, version string, configURL string) string {
	rootLive := fmtRootLive(baseURL, version)
	return fmt.Sprintf("ip=dhcp net.ifnames=1 rd.cos.disable rd.noverifyssl console=tty1 harvester.install.automatic=true root=live:%s harvester.install.config_url=%s", rootLive, configURL)
}

func fmtResp(cluster config.Cluster, node config.Node) (gin.H, error) {
	value := config.Value{
		BaseURL: config.Conf.BaseURL,
		OS:      config.Conf.OS,
		Cluster: cluster,
		Node:    node,
	}
	harvesterConfigContent, err := config.Render("harvester_config.tmpl", value)
	if err != nil {
		return nil, err
	}

	configFileName := fmtConfigFileName(node.Mac)
	if err = util.CreateFileWithContent(configFileName, harvesterConfigContent); err != nil {
		return nil, err
	}

	configURL := fmtConfigURL(configFileName)
	return gin.H{
		"kernel":  fmtKernel(config.Conf.BaseURL, cluster.Version),
		"initrd":  []string{fmtInitrd(config.Conf.BaseURL, cluster.Version)},
		"cmdline": fmtCmdline(config.Conf.BaseURL, cluster.Version, configURL),
	}, nil
}
