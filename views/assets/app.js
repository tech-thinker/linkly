/*
 * app.js
 */
let darkMode = localStorage.getItem("darkMode");
const themeToggle = document.getElementById("themeButton");
// Applying dark mode
const enableDarkMode = () => {
  document.body.classList.add("dark-mode");
  localStorage.setItem("darkMode", "enabled");
  themeToggle.innerHTML = `
        <svg
              style="
                width: 2rem;
                height: 2rem;
                color: #eceff1;
                display: inline-block;
              "
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokelinecap="round"
                strokelinejoin="round"
                strokewidth="2"
                d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"
              />
            </svg>`;
};
// Applying light mode
const disableDarkMode = () => {
  document.body.classList.remove("dark-mode");
  localStorage.setItem("darkMode", null);
  themeToggle.innerHTML = `
        <svg
              style="
                width: 2rem;
                height: 2rem;
                color: #ffa726;
                display: inline-block;
              "
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokelinecap="round"
                strokelinejoin="round"
                strokewidth="2"
                d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z"
              />
            </svg>`;
};
// Checking if dark mode is enabled
if (darkMode === "enabled") {
  enableDarkMode();
} else {
  disableDarkMode();
}

// Add event listener to toggle button
themeToggle.addEventListener("click", () => {
  if (darkMode !== "enabled") {
    enableDarkMode();
  } else {
    disableDarkMode();
  }
});

// Only works in input form
// function copytoClipboard(id) {
//   var copyText = document.getElementById(id);
//   copyText.select();
//   copyText.setSelectionRange(0, 99999);
//   navigator.clipboard.writeText(copyText.value);
// }
