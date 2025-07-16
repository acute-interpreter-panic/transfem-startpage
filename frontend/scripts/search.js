console.log("adding features to search...");

const form = document.getElementById("search-form");
const input = document.getElementById("search-input");

// https://stackoverflow.com/a/3809435/16804841
const expression =
  /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)/gi;
const urlRegex = new RegExp(expression);

const searchEngines = {
  g: {
    action: "https://www.google.com/search",
    name: "q",
  },
  d: {
    action: "https://duckduckgo.com/",
    name: "q",
  },
  y: {
    action: "https://www.youtube.com/results",
    name: "search_query",
  },
  ya: {
    action: "https://yandex.com/search/",
    name: "text",
  },
  lure: {
    action: "https://lure.sh/pkgs",
    name: "q",
  },
};

const translationPrefixes = ["t", "translation"];

function getDeepLUrl(s) {
  const parts = s.split("-");
  if (parts.length != 3) {
    return undefined;
  }

  return `https://www.deepl.com/en/translator?/#${encodeURIComponent(
    parts[0].trim()
  )}/${encodeURIComponent(parts[1].trim())}/${encodeURIComponent(
    parts[2].trim()
  )}`;
}

form.addEventListener("submit", (event) => {
  event.preventDefault();

  s = input.value;

  // check if url
  if (s.match(urlRegex)) {
    window.open(s, "_self");
    return;
  }

  // deepl translations
  let doTranslation = false;
  for (const value of translationPrefixes) {
    const prefix = `!${value} `;
    if (s.startsWith(prefix)) {
      doTranslation = true;
      s = s.slice(prefix.length); // Remove the !{key} prefix
      break;
    }
  }

  if (doTranslation) {
    const url = getDeepLUrl(s);
    if (url) {
      window.open(url.toString(), "_self");
      return;
    }
  }

  // Check if the string starts with ! followed by a key from searchEngines
  let selectedEngine = {
    action: form.getAttribute("action"),
    name: input.getAttribute("name"),
  };

  for (const [key, value] of Object.entries(searchEngines)) {
    const prefix = `!${key} `;
    if (s.startsWith(prefix)) {
      selectedEngine = value;
      s = s.slice(prefix.length); // Remove the !{key} prefix
      break;
    }
  }

  const url = new URL(selectedEngine.action);
  url.searchParams.set(selectedEngine.name, s.trim());
  window.open(url.toString(), "_self");
});
