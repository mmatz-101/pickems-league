import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import Login from "./routes/login";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import ErrorPageNothingFound from "./routes/error-page";
import { MantineProvider } from "@mantine/core";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPageNothingFound />,
  },
  {
    path: "/login",
    element: <Login />,
  },
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <MantineProvider
      theme={{
        colorScheme: "dark",
        headings: {
          fontFamily: "Roboto, sans-serif",
        },
      }}
      withGlobalStyles
      withNormalizeCSS
    >
      <RouterProvider router={router} />
    </MantineProvider>
  </React.StrictMode>
);
