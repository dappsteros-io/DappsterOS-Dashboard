import { createAction, createAsyncThunk } from "@reduxjs/toolkit";

import { signin, signup, getme } from "./api";

export const signIn = createAsyncThunk<any, any>("user/signin", async (data) =>
  signin(data)
);

export const signUp = createAsyncThunk<any, any>("user/signUp", async (data) =>
  signup(data)
);

export const getMe = createAsyncThunk<any, any>("user/getMe", async () =>
  getme()
);
