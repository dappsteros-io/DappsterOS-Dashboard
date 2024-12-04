import { createAction, createAsyncThunk } from "@reduxjs/toolkit";

import { get_nodes, get_vms, start_vm, stop_vm } from "./api";

export const getNodes = createAsyncThunk<any, any>(
  "proxmox/getNodes",
  async (data) => get_nodes(data)
);

export const getVMs = createAsyncThunk<any, any>(
  "proxmox/getVMs",
  async (data) => get_vms(data)
);

export const startVM = createAsyncThunk<any, any>("proxmox/startVM", async (data) =>
  start_vm(data)
);

export const stopVM = createAsyncThunk<any, any>("proxmox/stopVM", async (data) =>
  stop_vm(data)
);
