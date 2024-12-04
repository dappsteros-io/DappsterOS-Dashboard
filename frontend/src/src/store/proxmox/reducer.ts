import { createReducer } from "@reduxjs/toolkit";
import { getNodes, getVMs } from "./actions";
import { NodeData, VMData } from "@/types/proxmox";
const initalState: { nodes: NodeData[]; vms: VMData[]; error: any } = {
  nodes: [],
  vms: [],
  error: {},
};
export const proxmoxReducer = createReducer(initalState, (builder) => {
  builder
    .addCase(getNodes.fulfilled, (state, action) => {
      state.nodes = action.payload?.data;
      state.error = {};
    })
    .addCase(getNodes.pending, (state, action) => {
      state.nodes = [];
      state.error = {};
    })

    .addCase(getNodes.rejected, (state, action) => {
      state.nodes = [];
      state.error = action.error;
    })
    .addCase(getVMs.fulfilled, (state, action) => {
      state.vms = action.payload?.data;
      state.error = {};
    })
    .addCase(getVMs.pending, (state, action) => {
      state.vms = [];
      state.error = {};
    })

    .addCase(getVMs.rejected, (state, action) => {
      state.vms = [];
      state.error = action.error;
    });
});
