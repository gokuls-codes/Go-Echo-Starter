document.addEventListener("DOMContentLoaded", (event) => {
  document.body.addEventListener("htmx:beforeSwap", (evt) => {
    if (evt.detail.xhr.status === 422) {
      evt.detail.shouldSwap = true;
      evt.detail.isError = false;
    }
    if (evt.detail.xhr.status === 500) {
      evt.detail.shouldSwap = true;
      evt.detail.isError = true;
    }
    if (evt.detail.xhr.status === 409) {
      evt.detail.shouldSwap = true;
      evt.detail.isError = true;
    }
  });
});
