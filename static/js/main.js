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
    if (event.keyCode == 13) {
        button.click();
    }
});

function resetHistory() {
    localStorage.setItem("message-history", JSON.stringify([]));
    htmx.trigger(document.body, "reset-trigger");
}

function getMessageHistory() {
    const messageHistory = localStorage.getItem("message-history");
    let messages = []
    if (messageHistory !== null) {
        messages = JSON.parse(messageHistory);
    }
    return messages
}
