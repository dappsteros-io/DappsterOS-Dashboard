import api from "@/api/index";

export const get_nodes = (credentials: any) => {
  return api("proxmox/nodes", { method: "GET" });
};

export const get_vms = (credentials: any) => {
  return api("proxmox/vms", { method: "GET" });
};

export const get_vm = (vmid: any) => {
  return api(`proxmox/vms/${vmid}`, { method: "GET", params: { vmid } });
};
export const create_vm = (data: any) => {
  return api("proxmox/vms", { method: "POST", data });
};
export const start_vm = (vmid: number) => {
  return api(`proxmox/vms/${vmid}/start`, { method: "POST" });
};
export const install_dappster = (vmid: number) => {
  return api(`proxmox/vms/${vmid}/install`, { method: "POST" });
};

export const check_dappster = (vmid: number) => {
  return api(`proxmox/vms/${vmid}/status`, { method: "POST" });
};

export const stop_vm = (vmid: number) => {
  return api(`proxmox/vms/${vmid}/stop`, { method: "POST" });
};

export const delete_vm = (vmid: number) => {
  return api(`proxmox/vms/${vmid}`, { method: "DELETE" });
};
