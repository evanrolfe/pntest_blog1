# Route all outgoing traffic to mitmproxy on port $PROXY_PORT, but dont apply iptables rules to mitmproxyuser
# iptables -t nat -A OUTPUT -m owner ! --uid-owner mitmproxyuser -p tcp --dport 1:65535 -j REDIRECT --to-port $PROXY_PORT

iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 1:65535 -j REDIRECT --to-port $PROXY_PORT

# Prevent an infinite loop by running mitmproxy as the mitmprpoxyuser whos traffic is not redirected by
su mitmproxyuser -c "mitmweb -p $PROXY_PORT --web-port $WEB_UI_PORT --web-host 0.0.0.0 --mode transparent"
