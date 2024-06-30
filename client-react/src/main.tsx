import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import { createBrowserRouter, Outlet, RouterProvider } from "react-router-dom";
import Teams from "./components/Teams.tsx";
import { ThemeProvider } from "./components/theme-provider.tsx";
import Header from "./components/Header.tsx";
import App from "./App.tsx";
import Positions from "./components/Positions.tsx";
import Nations from "./components/Nations.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <>
        <Header />
        <Outlet />
      </>
    ),
    children: [
      {
        path: "",
        element: <App />,
      },
      {
        path: "teams",
        element: <Teams />,
      },
      {
        path: "nations",
        element: <Nations />,
      },
      {
        path: "positions",
        element: <Positions />,
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <RouterProvider router={router} />
    </ThemeProvider>
  </React.StrictMode>,
);
