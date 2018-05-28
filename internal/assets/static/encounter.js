function rollInitiative() {
    // RAW: Needs dialog to edit state to pass to backend.
    post("/cmds/rollInitiative", function(xhttp) {
        if (xhttp.status == 200) {
            document.getElementById("content").innerHTML = xhttp.responseText;
        }
    }, JSON.stringify({
        "id" : 0
    }));
}

function globalOptions() {
    // RAW: Needs dialog to edit state to pass to backend.
    post("/cmds/globalOptions", function(xhttp) {
        if (xhttp.status == 200) {
            document.getElementById("content").innerHTML = xhttp.responseText;
        }
    }, JSON.stringify({
        "id" : 0
    }));
}

function newCombatant() {
    post("/cmds/newCombatant", function(xhttp) {
        if (xhttp.status == 200) {
            document.getElementById("content").innerHTML = xhttp.responseText;
        }
    });
}

function deleteAllEnemies() {
    simpleModal({
        content: "Delete all enemy combatants?",
        buttons: [
            {
                title: "Delete",
                onclick: function() {
                    post("/cmds/deleteAllEnemies", function(xhttp) {
                        if (xhttp.status == 200) {
                            document.getElementById("content").innerHTML = xhttp.responseText;
                        }
                        closeSimpleModal();
                    });
                },
            }
        ]
    });
}

function adjustHP(id) {
    // RAW: Needs dialog to edit state to pass to backend.
    post("/cmds/adjustHP", function(xhttp) {
        if (xhttp.status == 200) {
            document.getElementById("content").innerHTML = xhttp.responseText;
        }
    }, JSON.stringify({
        "id" : id
    }));
}

function editCombatant(id) {
    // RAW: Needs dialog to edit state to pass to backend.
    post("/cmds/editCombatant", function(xhttp) {
        if (xhttp.status == 200) {
            document.getElementById("content").innerHTML = xhttp.responseText;
        }
    }, JSON.stringify({
        "id" : id
    }));
}

function duplicateCombatant(id) {
    sendSimpleCommand("duplicateCombatant", id);
}

function makeCurrentCombatant(id) {
    sendSimpleCommand("makeCurrentCombatant", id);
}

function deleteCombatant(name, id) {
    simpleModal({
        content: "Delete " + name + "?",
        buttons: [
            {
                title: "Delete",
                onclick: function() {
                    post("/cmds/deleteCombatant", function(xhttp) {
                        if (xhttp.status == 200) {
                            document.getElementById("content").innerHTML = xhttp.responseText;
                        }
                        closeSimpleModal();
                    }, JSON.stringify({ "id" : id }));
                },
            }
        ]
    });
}

function addNote(id, index) {
    // RAW: Needs dialog to get initial state to pass to backend.
    post("/cmds/addNote", function(xhttp) {
        if (xhttp.status == 200) {
            document.getElementById("content").innerHTML = xhttp.responseText;
        }
    }, JSON.stringify({
        "id" : id,
        "index" : index
    }));
}

function editNote(id, index) {
    // RAW: Needs dialog to edit state to pass to backend.
    post("/cmds/editNote", function(xhttp) {
        if (xhttp.status == 200) {
            document.getElementById("content").innerHTML = xhttp.responseText;
        }
    }, JSON.stringify({
        "id" : id,
        "index" : index
    }));
}

function deleteNote(id, index) {
    post("/cmds/deleteNote", function(xhttp) {
        if (xhttp.status == 200) {
            document.getElementById("content").innerHTML = xhttp.responseText;
        }
    }, JSON.stringify({
        "id" : id,
        "index" : index
    }));
}

function nextTurn() {
    post("/cmds/nextTurn", function(xhttp) {
        if (xhttp.status == 200) {
            document.getElementById("content").innerHTML = xhttp.responseText;
            var text;
            var round = xhttp.getResponseHeader("round");
            if (round < 1) {
                text = "Awaiting Initiative";
            } else {
                text = "Round " + round;
            }
            document.getElementById("turn").innerText = text;
            if (xhttp.getResponseHeader("new_round") == "true") {
                window.speechSynthesis.cancel();
                window.speechSynthesis.speak(new SpeechSynthesisUtterance("start of " + text));
            }
        }
    });
}

function sendSimpleCommand(cmd, id) {
    post("/cmds/" + cmd, function(xhttp) {
        if (xhttp.status == 200) {
            document.getElementById("content").innerHTML = xhttp.responseText;
        }
    }, JSON.stringify({ "id" : id }));
}

function simpleModal(options) {
    var dialog = document.getElementById("simple-modal-dialog");
    dialog.querySelector(".modal-content").innerHTML = options.content;
    var buttons = dialog.querySelector(".modal-buttons");
    buttons.innerHTML = "";
    var i;
    var needAutofocus = true;
    var len = options.buttons.length;
    for (i = 0; i < len; i++) {
        var button = document.createElement("button");
        button.onclick = options.buttons[i].onclick;
        if (options.buttons[i].autofocus) {
            button.autofocus = true;
        }
        button.appendChild(document.createTextNode(options.buttons[i].title));
        buttons.appendChild(button);
    }
    var cancel = document.createElement("button");
    cancel.onclick = closeSimpleModal;
    if (needAutofocus) {
        cancel.autofocus = true;
    }
    cancel.appendChild(document.createTextNode("Cancel"));
    buttons.appendChild(cancel);
    document.getElementById("modal-overlay").classList.remove("closed");
    dialog.classList.remove("closed");
}

function closeSimpleModal() {
    document.getElementById("modal-overlay").classList.add("closed");
    document.getElementById("simple-modal-dialog").classList.add("closed");
}

function post(url, callback, data) {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4) {
            callback(xhttp);
        }
    };
    xhttp.open("POST", url, true);
    xhttp.setRequestHeader("Content-type", "application/json");
    xhttp.send(data);
}

function handleGlobalKeyDown(event) {
    switch (event.code) {
        case "KeyN":
            nextTurn();
            break;
    }
}