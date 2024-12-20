import { createAction, createAsyncThunk } from "@reduxjs/toolkit";

import {
  create_vm,
  get_nodes,
  get_vms,
  get_vm,
  start_vm,
  stop_vm,
  delete_vm,
  install_dappster,
  check_dappster,
} from "./api";

export const getNodes = createAsyncThunk<any, any>(
  "proxmox/getNodes",
  async (data) => get_nodes(data)
);

export const getVMs = createAsyncThunk<any, any>(
  "proxmox/getVMs",
  async (data) => get_vms(data)
);

export const getVM = createAsyncThunk<any, any>("proxmox/getVM", async (data) =>
  get_vm(data)
);
export const createVM = createAsyncThunk<any, any>(
  "proxmox/createVM",
  async (data) => create_vm(data)
);

export const startVM = createAsyncThunk<any, any>(
  "proxmox/startVM",
  async (data) => start_vm(data)
);

export const installDappster = createAsyncThunk<any, any>(
  "proxmox/installDappster",
  async (data) => install_dappster(data)
);

export const checkDappster = createAsyncThunk<any, any>(
  "proxmox/checkDappster",
  async (data) => check_dappster(data)
);

export const stopVM = createAsyncThunk<any, any>(
  "proxmox/stopVM",
  async (data) => stop_vm(data)
);

export const deleteVM = createAsyncThunk<any, any>(
  "proxmox/deleteVM",
  async (data) => delete_vm(data)
);
