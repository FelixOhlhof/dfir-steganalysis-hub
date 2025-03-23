#!/bin/sh

cat <<EOF > /usr/share/nginx/html/env-config.js
window._env_ = {
  REACT_APP_RESTGW_URL: "$REACT_APP_RESTGW_URL",
};
EOF

nginx -g "daemon off;"
