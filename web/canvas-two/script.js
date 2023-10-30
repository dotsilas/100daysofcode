// SELECT TO ELEMENTS

// HTML Selectors
// const body = document.getElementByTagName("body");
// const paragraph = documet.getElementByClassName("paragraph");
let h1Element = document.getElementById("title-text");
let scriptELement = document.querySelector("script");

// CSS Selectors
// const allParagraphInDiv = document.querySelectorAll("div p");
// const paragraph = document.querySelector("div p");

// Create elements
let oElement = document.createElement("ol");
let li1 = document.createElement("li");
let text = document.createTextNode("First Member of list");

// ---------- insert element to html -------------
// document.body.insertBefore(oElement, scriptELement);

// ---------- insert element relative to another element ----------
h1Element.insertAdjacentElement("afterend", oElement);

// ---------- Append elements ----------
li1.appendChild(text);
oElement.appendChild(li1);

// More simple way to create
let orderedList = document.createElement("ol");
oElement.insertAdjacentElement("afterend", orderedList);

let languages = [
  "Golang",
  "C-sharp",
  "Zig",
  "Python",
  "Javascript",
  "Typescript"
];

languages.forEach((language) => {
  let li = document.createElement("li");
  li.innerHTML = language;
  orderedList.appendChild(li);
});

// delete and move elements
let unorderedList = document.createElement("ul"); // new list
orderedList.insertAdjacentElement("afterend", unorderedList); // add after ordered

// lastli == last item in OrderedList
let lastLi = orderedList.lastChild;
// delete lastli
lastLi.remove();
// add lastli to unorderedList
unorderedList.appendChild(lastLi);

// CHANGE ELEMENTS CLASSES
const body = document.getElementsByTagName("body");


