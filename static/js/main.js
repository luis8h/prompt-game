document.body.addEventListener("resetChatHistory", resetHistory);

document.body.addEventListener("htmx:configRequest", function(evt) {
    if (evt.detail.target.id === "chat") {
        const messageHistory = localStorage.getItem("message-history") || "[]";
        evt.detail.parameters.messages = JSON.stringify(JSON.parse(messageHistory));
    }
});

document.body.addEventListener("htmx:afterRequest", function(evt) {
    if (evt.detail.pathInfo.requestPath === "/message/user") {
        const promptInput = document.querySelector("#prompt-input");

        let messages = getMessageHistory();

        const userMessage = promptInput.value
        messages.push({ role: "user", content: userMessage })
        promptInput.value = "";
        adjustHeight(promptInput)

        fetch("/prompt", {
            method: "POST",
            body: JSON.stringify({ messages: messages })
        })
            .then(response => response.json())
            .then(data => {
                htmx.ajax("POST", "/message/assistant", { values: { message: data.answer }, swap: "beforeend", target: "#chat-history" })
                messages.push({ role: "assistant", content: data.answer })
                localStorage.setItem("message-history", JSON.stringify(messages))
            })
    }
});

document.querySelector("#prompt-input").addEventListener("keydown", function(event) {
    const button = document.querySelector("#send-button");
    if (event.keyCode === 13 && !event.shiftKey) {
        event.preventDefault();
        button.click();
        return
    }
    adjustHeight(event.target);
});

document.querySelector("#prompt-input").addEventListener("input", function(event) {
    adjustHeight(event.target); // Adjust height on input as usual
});

function getMessageHistory() {
    const messageHistory = localStorage.getItem("message-history");
    let messages = []
    if (messageHistory !== null) {
        messages = JSON.parse(messageHistory);
    }
    return messages
}

function resetHistory() {
    localStorage.setItem("message-history", JSON.stringify([]));
}

window.scrollChatToTop = function() {
    const simpleBar = SimpleBar.instances.get(document.getElementById("chat"));
    if (simpleBar) {
        const scrollElement = simpleBar.getScrollElement();
        scrollElement.scrollTop = scrollElement.scrollHeight;
    }
};

window.onResetClick = function() {
    resetHistory()
    htmx.trigger(document.body, "reset-trigger");
};

function adjustHeight(textarea) {
    textarea.style.height = 'auto'; // Reset height to recalculate
    textarea.style.height = textarea.scrollHeight + 'px'; // Set height to fit content
}


// fix simple bar not working with htmx because the content is swapped out and not reinitialized
document.addEventListener("htmx:afterSwap", function(event) {
    const swappedContent = event.detail.target;
    if (swappedContent.hasAttribute("data-simplebar")) {
        reinitializeSimplebar(swappedContent);
    }
    swappedContent.querySelectorAll("[data-simplebar]").forEach((element) => {
        reinitializeSimplebar(element);
    });
});
function reinitializeSimplebar(element) {
    if (SimpleBar.instances.has(element)) {
        SimpleBar.instances.get(element).unMount();
    }
    new SimpleBar(element, { autoHide: false });
}
