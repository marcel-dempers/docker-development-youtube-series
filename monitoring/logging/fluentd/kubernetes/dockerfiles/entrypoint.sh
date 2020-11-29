#!/usr/bin/env sh

exec fluentd -c /fluentd/etc/${FLUENTD_CONF} -p /fluentd/plugins