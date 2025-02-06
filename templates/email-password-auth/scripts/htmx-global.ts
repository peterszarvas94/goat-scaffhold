// this file is needed so I can use htmx types
import * as h from "./htmx/htmx";

declare global {
  const htmx: typeof h;
}

export {};
