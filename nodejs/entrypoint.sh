#!/bin/sh
export TOKEN=$(curl -X PUT http://169.254.169.254/latest/api/token -H "X-aws-ec2-metadata-token-ttl-seconds: 21600")
export MW_AGENT_SERVICE=$(curl http://169.254.169.254/latest/meta-data/local-hostname -H "X-aws-ec2-metadata-token: $TOKEN")
echo $MW_AGENT_SERVICE
node app.js
