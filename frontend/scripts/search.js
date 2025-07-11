console.log("adding features to search...");

const form = document.getElementById("search-form");
const input = document.getElementById("search-input");

const searchEngines = {
    "g": {
        action: "https://www.google.com/search/",
        name: "q",
    },
    "d": {
        action: "https://duckduckgo.com/",
        name: "q",
    },
    "y": {
        action: "https://yandex.com/search/",
        name: "text",
    },
    "lure": {
        action: "https://lure.sh/pkgs",
        name: "q",
    },
};

// https://stackoverflow.com/a/3809435/16804841
const expression = /[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)/gi;
const urlRegex = new RegExp(expression);

form.addEventListener("submit", event => {
    s = input.value;

    // check if url
    if (s.match(urlRegex)) {
        event.preventDefault();
        window.open(s, "_self");
        return
    }
});
