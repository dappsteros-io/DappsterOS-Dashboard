import api from "@/api/index";

export const signin = (credentials: any) => {
  return api("auth/signin", { method: "POST", data: credentials });
};

export const signup = (credentials: any) => {
  return api("auth/signup", { method: "POST", data: credentials });
};

export const getme = () => {
  return api("auth/me", { method: "GET" });
};
