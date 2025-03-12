import htmx from "htmx.org";

htmx.logAll();

/** @type {Partial<import("htmx.org").HtmxExtension>} */
const swapAll = {
	onEvent: function (name, event) {
		if (name === "htmx:beforeSwap") {
			/** @type {CustomEvent} */ (event).detail.shouldSwap = true;
		}
		return true;
	},
};

htmx.defineExtension("swap-all", swapAll);
