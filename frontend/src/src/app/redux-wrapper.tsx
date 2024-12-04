"use client";

import { Provider } from "react-redux";
import { initializeStore } from "@/store";

interface ReduxProviderProps {
  children: React.ReactNode;
}

const ReduxProvider: React.FC<ReduxProviderProps> = ({ children }) => {
  const store = initializeStore(); 
  return <Provider store={store}>{children}</Provider>;
};

export default ReduxProvider;
