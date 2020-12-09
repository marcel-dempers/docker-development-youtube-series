#
# Fluentd
#
#    Licensed under the Apache License, Version 2.0 (the "License");
#    you may not use this file except in compliance with the License.
#    You may obtain a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
#    Unless required by applicable law or agreed to in writing, software
#    distributed under the License is distributed on an "AS IS" BASIS,
#    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#    See the License for the specific language governing permissions and
#    limitations under the License.
#

# The following Fluentd parser plugin, aims to simplify the parsing of multiline
# logs found in Kubernetes nodes. Since many log files shared the same format and
# in order to simplify the configuration, this plugin provides a 'kubernetes' format
# parser (built on top of MultilineParser).
#
# When tailing files, this 'kubernetes' format should be applied to the following
# log file sources:
#
#  - /var/log/kubelet.log
#  - /var/log/kube-proxy.log
#  - /var/log/kube-apiserver.log
#  - /var/log/kube-controller-manager.log
#  - /var/log/kube-scheduler.log
#  - /var/log/rescheduler.log
#  - /var/log/glbc.log
#  - /var/log/cluster-autoscaler.log
#
# Usage:
#
# ---- fluentd.conf ----
#
# <source>
#   @type tail
#   path ./kubelet.log
#   read_from_head yes
#   tag kubelet
#   <parse>
#     @type kubernetes
#   </parse>
# </source>
#
# ----   EOF       ---

require 'fluent/plugin/parser_regexp'

module Fluent
  module Plugin
    class KubernetesParser < RegexpParser
      Fluent::Plugin.register_parser("kubernetes", self)

      CONF_FORMAT_FIRSTLINE = %q{/^\w\d{4}/}
      CONF_FORMAT1 = %q{/^(?<severity>\w)(?<time>\d{4} [^\s]*)\s+(?<pid>\d+)\s+(?<source>[^ \]]+)\] (?<message>.*)/m}
      CONF_TIME_FORMAT = "%m%d %H:%M:%S.%N"

      def configure(conf)
        conf['expression'] = CONF_FORMAT1
        conf['time_format'] = CONF_TIME_FORMAT
        super
      end
    end
  end
end