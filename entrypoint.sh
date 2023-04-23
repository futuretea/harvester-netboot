#!/bin/bash
set -e

exec tini -- app "${@}"
