#!/usr/bin/env sh

# The MIT License
#
#  Copyright (c) 2015, CloudBees, Inc.
#
#  Permission is hereby granted, free of charge, to any person obtaining a copy
#  of this software and associated documentation files (the "Software"), to deal
#  in the Software without restriction, including without limitation the rights
#  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
#  copies of the Software, and to permit persons to whom the Software is
#  furnished to do so, subject to the following conditions:
#
#  The above copyright notice and this permission notice shall be included in
#  all copies or substantial portions of the Software.
#
#  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
#  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
#  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
#  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
#  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
#  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
#  THE SOFTWARE.

# Usage jenkins-slave.sh [options] -url http://jenkins [SECRET] [AGENT_NAME]
# Optional environment variables :
# * JENKINS_TUNNEL : HOST:PORT for a tunnel to route TCP traffic to jenkins host, when jenkins can't be directly accessed over network
# * JENKINS_URL : alternate jenkins URL
# * JENKINS_SECRET : agent secret, if not set as an argument
# * JENKINS_AGENT_NAME : agent name, if not set as an argument
# * JENKINS_AGENT_WORKDIR : agent work directory, if not set by optional parameter -workDir

if [ $# -eq 1 ]; then

	# if `docker run` only has one arguments, we assume user is running alternate command like `bash` to inspect the image
	exec "$@"

else

	# if -tunnel is not provided try env vars
	case "$@" in
		*"-tunnel "*) ;;
		*)
		if [ ! -z "$JENKINS_TUNNEL" ]; then
			TUNNEL="-tunnel $JENKINS_TUNNEL"
		fi ;;
	esac

	# if -workDir is not provided try env vars
	if [ ! -z "$JENKINS_AGENT_WORKDIR" ]; then
		case "$@" in
			*"-workDir"*) echo "Warning: Work directory is defined twice in command-line arguments and the environment variable" ;;
			*)
			WORKDIR="-workDir $JENKINS_AGENT_WORKDIR" ;;
		esac
	fi

	if [ -n "$JENKINS_URL" ]; then
		URL="-url $JENKINS_URL"
	fi

	if [ -n "$JENKINS_NAME" ]; then
		JENKINS_AGENT_NAME="$JENKINS_NAME"
	fi  

	if [ -z "$JNLP_PROTOCOL_OPTS" ]; then
		echo "Warning: JnlpProtocol3 is disabled by default, use JNLP_PROTOCOL_OPTS to alter the behavior"
		JNLP_PROTOCOL_OPTS="-Dorg.jenkinsci.remoting.engine.JnlpProtocol3.disabled=true"
	fi

	# If both required options are defined, do not pass the parameters
	OPT_JENKINS_SECRET=""
	if [ -n "$JENKINS_SECRET" ]; then
		case "$@" in
			*"${JENKINS_SECRET}"*) echo "Warning: SECRET is defined twice in command-line arguments and the environment variable" ;;
			*)
			OPT_JENKINS_SECRET="${JENKINS_SECRET}" ;;
		esac
	fi
	
	OPT_JENKINS_AGENT_NAME=""
	if [ -n "$JENKINS_AGENT_NAME" ]; then
		case "$@" in
			*"${JENKINS_AGENT_NAME}"*) echo "Warning: AGENT_NAME is defined twice in command-line arguments and the environment variable" ;;
			*)
			OPT_JENKINS_AGENT_NAME="${JENKINS_AGENT_NAME}" ;;
		esac
	fi

	#TODO: Handle the case when the command-line and Environment variable contain different values.
	#It is fine it blows up for now since it should lead to an error anyway.

	exec java $JAVA_OPTS $JNLP_PROTOCOL_OPTS -cp /usr/share/jenkins/slave.jar hudson.remoting.jnlp.Main -headless $TUNNEL $URL $WORKDIR $OPT_JENKINS_SECRET $OPT_JENKINS_AGENT_NAME "$@"
fi