/*
 * client.js
 */

let baseAPI = "http://localhost:8080";
baseAPI = window.location.origin;
let apiURL = baseAPI + "/api/links";

// icons
let clipboardIcon = `<svg
    xmlns="http://www.w3.org/2000/svg"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    >
    <path
      d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"
    ></path>
    <rect x="8" y="2" width="8" height="4" rx="1" ry="1"></rect>
    </svg>`;
let tickIcon = `<svg
    xmlns="http://www.w3.org/2000/svg"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    >
      <path
        d="M12 0c6.623 0 12 5.377 12 12s-5.377 12-12 12-12-5.377-12-12 5.377-12 12-12zm0 1c6.071 0 11 4.929 11 11s-4.929 11-11 11-11-4.929-11-11 4.929-11 11-11zm7 7.457l-9.005 9.565-4.995-5.865.761-.649 4.271 5.016 8.24-8.752.728.685z"
      />
    </svg>`;

// getnewURL returns a new short url
function generateURL() {
  let url = document.getElementById("input_url").value;
  let output = document.getElementById("shorten_url");
  if (url == "" || url == null) {
    output.innerHTML = `
    <div>
      <h3>
        <span class="text-warning"> Please Enter a url </span>
      </h3>
    </div>`;
    setTimeout(() => {
      output.innerHTML = "";
    }, 3000);
    return;
  }
  url = trimURL(url);
  if (!validateURL(url)) {
    output.innerHTML = `
    <div>
      <h3>
        <span class="text-danger"> Please Enter a valid url </span>
      </h3>
    </div>`;
    setTimeout(() => {
      output.innerHTML = "";
    }, 3000);
    return;
  }
  let data = JSON.stringify({
    url: url,
  });
  fetch(apiURL, {
    // credentials: "same-origin",
    // mode: "same-origin",
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: data,
  })
    .then((resp) => {
      return resp.json();
    })
    .then((data) => {
      // console.log(data);
      output.innerHTML = `
      <div class="short_url">
        <h3>
          <a  id="output_url" href="${baseAPI}/${data.urls.short_url}" target="_blank">${baseAPI}/${data.urls.short_url}</a>
        </h3>
        <br/>
        <button id="clipboard" type="button" class="" onclick="copyToClipboard('output_url')">
        </button>
      </div>`;
      document.getElementById("clipboard").innerHTML = clipboardIcon;
      setTimeout(() => {
        output.innerHTML = "";
      }, 15000);
    })
    .catch((err) => {
      if (err === "server") return;
      console.log(err);
    });
  document.getElementById("input_url").value = "";
}

// velidateURL validates the url
function validateURL(url) {
  // string has containing dot then return true
  if (url.includes(".")) {
    return true;
  }
  return false;
}

// trimURL removes the http:// or https:// from the url and returns the url
function trimURL(url) {
  return url.replace(/^https?:\/\//, "").replace(/\/$/, "");
}

// clipboard toggle
let toggleClipboard = false;
// copyToClipboard copies the text to the clipboard
function copyToClipboard(elementId) {
  // Create a "hidden" input
  var aux = document.createElement("input");

  // Assign it the value of the specified element
  aux.setAttribute("value", document.getElementById(elementId).innerHTML);

  // Append it to the body
  document.body.appendChild(aux);

  // Highlight its content
  aux.select();

  // Copy the highlighted text
  document.execCommand("copy");

  // Remove it from the body
  document.body.removeChild(aux);

  // Toggle Icon
  clipButton = document.getElementById("clipboard");
  if (toggleClipboard == false) {
    clipButton.innerHTML = tickIcon;
    toggleClipboard = true;
  } else {
    clipButton.innerHTML = clipboardIcon;
    toggleClipboard = false;
  }
}

let stateHome = ``;
let toggleURL = false;
// viewURLs returns url list in home page
function viewURLs() {
  let baseAPI = window.location.origin;
  let content = document.getElementById("content_container");
  let urlPage = `
  <div>
  <h1><span>Linkly</span>: URLs</h1>
  <hr/>
  <div id="urls"></div>
  </div>
  `;
  if (toggleURL == false) {
    stateHome = content.innerHTML;
    content.innerHTML = urlPage;
    toggleURL = true;
  } else {
    content.innerHTML = stateHome;
    toggleURL = false;
    return;
  }
  fetch(baseAPI + "/api/links")
    .then((response) => response.json())
    .then((data) => {
      // console.log(data);
      for (let i = 0; i < data.urls.length; i++) {
        document.getElementById("urls").innerHTML += `
      <div class="url_element">
        <h2 class="">
        ${data.urls[i].url} -> <a href="/${data.urls[i].short_url}">
          ${baseAPI}/${data.urls[i].short_url}</a> , visits: ${data.urls[i].visits}
        </h2>
        <hr />
      </div>
        `;
      }
    });
}

// Add event listener to keypress
document.addEventListener("keypress", function (e) {
  if (e.keyCode === 13) {
    generateURL();
  }
});
