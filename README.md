# Unexpected (possibly bugged) behaviour of remote_address rate limit action

# demo.sh

./demo.sh will run docker-compose and curl and demonstrate the issue.

# Description

If Envoy is configured as a proxy with rate limiting using the `remote_address` action and both:

```
#...
                use_remote_address: true
                xff_num_trusted_hops: 0
#...
```

Then the `remote_action` action will not match a request which does not include a `x-forwarded-for` header, i.e.:

```
curl http://localhost:8888/test # this will not cause a descriptor to be generated and never gets rate limited
```

but a request with an `x-forwarded-for` header will:

```
curl -H 'x-forwarded-for: 1.1.1.1' http://localhost:8888/test # this will cause a descriptor to be generated
```

*and* in the later case the rate limit descriptor will use the downstream connection IP as the key (which is the correct behaviour!) and not the value passed in `x-forwarded-for`. 

I belive this is not the correct behaviour and in this set up the `remote_action` rate limit action should apply if no `x-forwarded-for` is sent to the proxy and should use the
downstream connection IP as the key.
