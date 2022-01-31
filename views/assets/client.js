/*
 * client.js
 */

let baseAPI = "http://localhost:8080";
baseAPI = window.location.origin;
let apiURL = baseAPI + "/api/links";

function generateURL() {
  let url = document.getElementById("input_url").value;
  let output = document.getElementById("shorten_url");
  if (url == "" || url == null) {
    output.innerHTML = `
    <div>
      <h3>
        <span class="text-danger"> Please Enter a valid url </span>
      </h3>
    </div>`;
    return;
  }
  url = trimURL(url);
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
      console.log(data);
      output.innerHTML = `
      <div class="short_url">
        <h3>
          <a  id="output_url" href="${baseAPI}/${data.urls.short_url}" target="_blank">${baseAPI}/${data.urls.short_url}</a>
        </h3>
        <br/>
        <button id="clipboard" type="button" class="" onclick="copyToClipboard('output_url')">
          <svg
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
          </svg>
        </button>
      </div>`;
    })
    .catch((err) => {
      if (err === "server") return;
      console.log(err);
    });
  document.getElementById("input_url").value = "";
}

function trimURL(url) {
  return url.replace(/^https?:\/\//, "").replace(/\/$/, "");
}

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
}

function viewURLs() {
  let baseAPI = window.location.origin;
  let content = document.getElementById("content_container");
  content.innerHTML = `
  <div>
  <h1><span>Linkly</span>: URLs</h1>
  <hr/>
  <div id="urls"></div>
  </div>
  `;
  fetch(baseAPI + "/api/links")
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
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
