import { configureStore, combineReducers } from "@reduxjs/toolkit";
import { userReducer } from "./user/reducer";

import { createWrapper, HYDRATE } from "next-redux-wrapper";
import { proxmoxReducer } from "./proxmox/reducer";

const rootReducer = combineReducers({
  user: userReducer,
  proxmox: proxmoxReducer,
});

export const makeStore = () => {
  return configureStore({
    reducer: rootReducer,
    devTools: process.env.NODE_ENV !== "production",
  });
};

let storeInstance: ReturnType<typeof configureStore>;

// Singleton store function
export const initializeStore = (preloadedState?: RootState) => {
  if (!storeInstance) {
    storeInstance = configureStore({
      reducer: rootReducer,
      devTools: process.env.NODE_ENV !== "production",
      preloadedState,
    });
  }
  return storeInstance;
};

// Access the existing store instance
export const getStore = () => storeInstance;

// Infer the type of makeStore
export type AppStore = ReturnType<typeof makeStore>;
// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<AppStore["getState"]>;
export type AppDispatch = AppStore["dispatch"];

export const wrapper = createWrapper(makeStore);
