document.body.addEventListener("htmx:configRequest", function(evt) {
    if (evt.detail.target.id === "chat") {
        const messageHistory = localStorage.getItem("message-history") || "[]";
        evt.detail.parameters.messages = JSON.stringify(JSON.parse(messageHistory));
    }
});

function highlightSubmit() {
    const submitButton = document.querySelector("#submit-button");
    submitButton.classList.add("animation-pulse");
    console.log("adding")
    setTimeout(() => {
        submitButton.classList.remove("animation-pulse");
        console.log("removing")
    }, 1500);
}

document.body.addEventListener("htmx:afterRequest", function(evt) {
    if (evt.detail.pathInfo.requestPath === "/message/user") {
        document.querySelector("#chat-loading").style.display = "block";
        document.querySelector("#send-button button").disabled = true
        scrollChatToTop();
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
                document.querySelector("#chat-loading").style.display = "none";
                document.querySelector("#send-button button").disabled = false
                htmx.ajax("POST", "/message/assistant", { values: { message: data.answer }, swap: "beforeend", target: "#chat-history" })
                    .then(() => {
                        scrollChatToTop();
                    });
                messages.push({ role: "assistant", content: data.answer })
                localStorage.setItem("message-history", JSON.stringify(messages))
                updateSubmit();
                updateSendButton();
                updateShowStrategy();
            })
    }
});

function updateSubmit() {
    let messages = getMessageHistory();
    const button = document.querySelector("#submit-button .button-1");
    const withStrategy = JSON.parse(document.getElementById("level-html").getAttribute("with-strategy"));
    const hasStrategy = JSON.parse(document.getElementById("level-html").getAttribute("has-strategy"));
    if (messages.length === 0 || (!withStrategy && hasStrategy)) {
        button.disabled = true;
    } else if (button.disabled === true) {
        button.disabled = false;
        highlightSubmit();
    }
}
window.updateSubmit = updateSubmit;

function updateSendButton() {
    const input = document.querySelector("#prompt-input");
    const button = document.querySelector("#send-button .button-2");
    const initialValue = input.value.trim();
    button.disabled = initialValue === "";
}
window.updateSendButton = updateSendButton;

function updateShowStrategy() {
    const messages = getMessageHistory();
    const strategy = document.querySelector("#strategy");
    if (!strategy) {
        return;
    }
    const withStrategy = JSON.parse(document.getElementById("level-html").getAttribute("with-strategy"));
    console.log(withStrategy)
    if (messages.length === 0 && !withStrategy) {
        strategy.classList.add("hidden");
    } else {
        strategy.classList.remove("hidden");
    }
}
window.updateShowStrategy = updateShowStrategy;

// enable/disable send button
document.querySelector("#prompt-input").addEventListener("input", function(event) {
    updateSendButton();
    adjustHeight(event.target);
});

// initial check when the page loads
document.addEventListener('DOMContentLoaded', function() {
    updateShowStrategy();
    updateSendButton();
    updateSubmit();
});


// check for (shift)enter press in input field
document.querySelector("#prompt-input").addEventListener("keydown", function(event) {
    const button = document.querySelector("#send-button");
    if (event.target.value.trim() === "") {
        return;
    }
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
    htmx.trigger(document.body, "reset-trigger");
    updateSubmit();
    updateShowStrategy();
}

function refreshLevel() {
    console.log("refreshing level")
    htmx.trigger(document.body, "refresh-level");
}
document.body.addEventListener("refreshLevel", refreshLevel);

function showInvalidAnswerPopup() {
    Toastify({
        text: "Sorry, the answer/strategy was not correct. Try again.",
        duration: 3000,
        close: true,
        gravity: "top",
        position: "center",
        backgroundColor: "#e53e3e",
    }).showToast();
}

function scrollChatToTop() {
    const simpleBar = SimpleBar.instances.get(document.getElementById("chat"));
    if (simpleBar) {
        const scrollElement = simpleBar.getScrollElement();
        scrollElement.scrollTo({
            top: scrollElement.scrollHeight,
            behavior: 'smooth'
        });
    }
};
window.scrollChatToTop = scrollChatToTop

window.onResetClick = resetHistory

document.body.addEventListener("resetChatHistory", resetHistory);
document.body.addEventListener("invalidAnswer", showInvalidAnswerPopup);
document.body.addEventListener("refreshShowStrategy", updateShowStrategy);
document.body.addEventListener("refreshSubmitButton", updateSubmit);

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
