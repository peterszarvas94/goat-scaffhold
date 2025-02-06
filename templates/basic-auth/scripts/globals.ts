// this file is needed so we can use types
import * as h from "./htmx/htmx";

declare global {
  const htmx: typeof h;
}

export {};
