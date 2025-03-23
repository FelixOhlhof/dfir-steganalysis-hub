const config = {
  restGwUrl: window._env_
    ? window._env_.REACT_APP_RESTGW_URL
    : process.env.REACT_APP_RESTGW_URL,
};

export default config;
