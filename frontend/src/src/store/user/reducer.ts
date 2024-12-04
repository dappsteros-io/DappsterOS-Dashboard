import { createReducer } from "@reduxjs/toolkit";
import { getMe, signIn, signUp } from "./actions";
import { UserData } from "@/types/user";
const initalState: { data: UserData; error: any } = {
  data: {},
  error: {},
};
export const userReducer = createReducer(initalState, (builder) => {
  builder

    .addCase(signIn.fulfilled, (state, action) => {
      state.data = action.payload?.data;
      localStorage.setItem("token", action.payload?.data.token);
      state.error = {};
    })
    .addCase(signIn.pending, (state, action) => {
      state.data = {};
      state.error = {};
    })

    .addCase(signIn.rejected, (state, action) => {
      state.data = {};
      state.error = action.error;
    })
    .addCase(signUp.fulfilled, (state, action) => {
      state.data = action.payload?.data;
    })
    .addCase(signUp.pending, (state, action) => {
      // state.data = []
    })

    .addCase(signUp.rejected, (state, action) => {
      // state.data = {};
    })
    .addCase(getMe.fulfilled, (state, action) => {
      state.data = action.payload?.data;
    })
    .addCase(getMe.pending, (state, action) => {
      // state.data = []
    })

    .addCase(getMe.rejected, (state, action) => {
      // state.data = {};
    });
});
