#!/bin/sh
#export TOKEN=$(curl -X PUT http://169.254.169.254/latest/api/token -H "X-aws-ec2-metadata-token-ttl-seconds: 21600")
#export MW_AGENT_SERVICE=$(curl http://169.254.169.254/latest/meta-data/local-ipv4 -H "X-aws-ec2-metadata-token: $TOKEN")
MW_AGENT_SERVICE=localhost
middleware-instrument \
--exporter_otlp_endpoint http://localhost:9319 \
--resource_attributes=project.name=python-webserver,mw.app.lang=python,runtime.metrics.python=true \
--service_name python-webserver \
app.py 
