#!/bin/bash

function cleanup {
  docker-compose -f docker-compose-remote-address-working.yml down > /dev/null 2>&1
  docker-compose -f docker-compose-remote-address-not-working.yml down > /dev/null 2>&1

}

trap cleanup EXIT

echo "In this first demo we will start a mock 'hello world' service, an envoy proxy in front of it, ratelimit and redis. The envoy proxy has not specified use_remote_address. A rate limit of 0 per second is configured and has 'failure_mode_deny: false', so if a request is rate limited it has matched a rate limit action and a successful request to the ratelimit cluster will have been made."
echo ""
read -p "Press enter to continue"
echo ""
echo "Starting services..."

docker-compose -f docker-compose-remote-address-working.yml up -d > /dev/null 2>&1

echo "Services up!"
echo ""
echo "Now let's curl the mock service via the proxy, and see that we are rate limited:"
echo ""
read -p "Press enter to continue"
echo ""
curl -I http://localhost:8888/test

echo ""
echo "We got a 429, this is expected."
read -p "Press enter to continue"
echo ""


echo "Tearing down services..."
docker-compose -f docker-compose-remote-address-working.yml down > /dev/null 2>&1
echo "Services stopped!"
echo ""


echo "Now we will restart the same services except now the proxy is configured with 'use_remote_address: true' and 'xff_num_trusted_hops: 0'"
echo ""
read -p "Press enter to continue"
echo ""
echo "Starting services..."

docker-compose -f docker-compose-remote-address-not-working.yml up -d > /dev/null 2>&1

echo "Services up!"
echo ""
echo "Now let's curl the mock service via the proxy, this time we will see that we are not rate limited:"
echo ""
read -p "Press enter to continue"
echo ""
curl -I http://localhost:8888/test

echo ""
echo "We got a 200, this is not expected (to me at least)"
echo ""
echo "Now let's curl the mock service via the proxy including a 'x-forwarded-for' header, this time we will see that we **are** rate limited:"
echo ""
read -p "Press enter to continue"
echo ""
curl -I -H "x-forwarded-for: 1.1.1.1" http://localhost:8888/test

echo ""
echo "Now we were rate limited. If you check the logs of the rate limiter you will see that 'remote_address' did not use the IP address from the 'x-forwarded-for' header but the machines local IP (this behaviour I expect!)"
echo ""
echo ""
echo "It's at least unexpected, and I think probably a bug, that in the second case the remote_address rate limit action did not apply without a x-forwarded-for' header being passed with the request. "
