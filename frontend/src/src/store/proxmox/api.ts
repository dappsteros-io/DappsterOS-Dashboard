import api from "@/api/index";

export const get_nodes = (credentials: any) => {
  return api("proxmox/nodes", { method: "GET" });
};

export const get_vms = (credentials: any) => {
  return api("proxmox/vms", { method: "GET" });
};

export const start_vm = (vmid: number) => {
  return api(`proxmox/vm/${vmid}/start`, { method: "POST" });
};

export const stop_vm = (vmid: number) => {
  return api(`proxmox/vm/${vmid}/stop`, { method: "POST" });
};