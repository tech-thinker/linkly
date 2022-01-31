/*
 * app.js
 */

// Store the theme
let darkTheme = localStorage.getItem("darkTheme");
const themeToggle = document.querySelector("#themeButton");
const bodyBackground = document.getElementById("#body");

// Apply Dark theme
const enableDark = () => {
  document.body.classList.add("dark-mode");
  localStorage.setItem("darkTheme", "enabled");
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

// Remove Dark theme
const disableDark = () => {
  document.body.classList.remove("dark-mode");
  localStorage.setItem("darkTheme", null);
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

//Toggle theme
if (darkTheme === "enabled") {
  enableDark();
  document.body.offsetHeight; // Trigger reflow to flush CSS changes
} else {
  disableDark();
}

themeToggle.addEventListener("click", () => {
  darkTheme = localStorage.getItem("darkTheme");
  if (darkTheme !== "enabled") {
    enableDark();
  } else {
    disableDark();
  }
});
