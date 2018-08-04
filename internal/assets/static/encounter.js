function rollInitiative() {
	post("/cmds/rollInitiative", function (xhttp) {
		if (xhttp.status == 200) {
			simpleModal({
				content: xhttp.responseText,
				wantAutoFocus: false,
				buttons: [{
					title: "Start Combat",
					autofocus: true,
					onclick: function () {
						var inputs = document.getElementById("fields").getElementsByTagName("input");
						var length = inputs.length;
						var inits = []
						for (var i = 0; i < length; i++) {
							inits.push({
								"id": inputs[i].name,
								"init": inputs[i].value
							})
						}
						post("/cmds/rollInitiative", function (xhttp) {
							if (xhttp.status == 200) {
								document.getElementById("board-area").innerHTML = xhttp.responseText;
								adjustRound(xhttp);
							}
							closeSimpleModal();
						}, JSON.stringify({
							"panel": false,
							"inits": inits
						}));
					},
				}]
			});
		}
	}, JSON.stringify({
		"panel": true
	}));
}

function globalOptions() {
	post("/cmds/globalOptions", function (xhttp) {
		if (xhttp.status == 200) {
			simpleModal({
				content: xhttp.responseText,
				wantAutoFocus: false,
				buttons: [{
					title: "Apply",
					autofocus: true,
					onclick: function () {
						var payload = {
							"panel": false
						};
						var inputs = document.getElementById("fields").getElementsByTagName("input");
						var length = inputs.length;
						for (var i = 0; i < length; i++) {
							payload[inputs[i].name] = inputs[i].value;
						}
						inputs = document.getElementById("fields").getElementsByTagName("select");
						length = inputs.length;
						for (var i = 0; i < length; i++) {
							payload[inputs[i].name] = inputs[i].value;
						}
						post("/cmds/globalOptions", function (xhttp) {
							if (xhttp.status == 200) {
								showEntity(selectedEntityElem(), true);
							}
							closeSimpleModal();
						}, JSON.stringify(payload));
					},
				}]
			});
		}
	}, JSON.stringify({
		"panel": true
	}));
}

function newCombatant(id) {
	var payload = {
		"panel": true
	};
	if (id !== undefined) {
		payload.panel = false;
		payload.based_on = id;
	}
	post("/cmds/newCombatant", function (xhttp) {
		if (xhttp.status == 200) {
			if (payload.panel) {
				simpleModal({
					content: xhttp.responseText,
					wantAutoFocus: false,
					buttons: [{
						title: "Add",
						autofocus: true,
						onclick: function () {
							var inputs = document.getElementById("fields").getElementsByTagName("input");
							var length = inputs.length;
							payload.panel = false;
							for (var i = 0; i < length; i++) {
								if (inputs[i].type == "checkbox") {
									payload[inputs[i].name] = inputs[i].checked;
								} else {
									payload[inputs[i].name] = inputs[i].value;
								}
							}
							post("/cmds/newCombatant", function (xhttp) {
								if (xhttp.status == 200) {
									document.getElementById("board-area").innerHTML = xhttp.responseText;
								}
								closeSimpleModal();
							}, JSON.stringify(payload));
						},
					}]
				});
			} else {
				document.getElementById("board-area").innerHTML = xhttp.responseText;
			}
		}
	}, JSON.stringify(payload));
}

function editCombatant(id) {
	post("/cmds/editCombatant", function (xhttp) {
		if (xhttp.status == 200) {
			simpleModal({
				content: xhttp.responseText,
				wantAutoFocus: false,
				buttons: [{
					title: "Change",
					autofocus: true,
					onclick: function () {
						var inputs = document.getElementById("fields").getElementsByTagName("input");
						var length = inputs.length;
						var payload = {
							"id": id,
							"panel": false
						}
						for (var i = 0; i < length; i++) {
							if (inputs[i].type == "checkbox") {
								payload[inputs[i].name] = inputs[i].checked;
							} else {
								payload[inputs[i].name] = inputs[i].value;
							}
						}
						post("/cmds/editCombatant", function (xhttp) {
							if (xhttp.status == 200) {
								document.getElementById("board-area").innerHTML = xhttp.responseText;
							}
							closeSimpleModal();
						}, JSON.stringify(payload));
					},
				}]
			});
		}
	}, JSON.stringify({
		"id": id,
		"panel": true
	}));
}

function deleteAllEnemies() {
	simpleModal({
		content: "Delete all enemy combatants?",
		buttons: [{
			title: "Delete",
			onclick: function () {
				post("/cmds/deleteAllEnemies", function (xhttp) {
					if (xhttp.status == 200) {
						document.getElementById("board-area").innerHTML = xhttp.responseText;
						adjustRound(xhttp);
					}
					closeSimpleModal();
				});
			},
		}]
	});
}

function adjustHP(id) {
	post("/cmds/adjustHP", function (xhttp) {
		if (xhttp.status == 200) {
			simpleModal({
				content: xhttp.responseText,
				wantAutoFocus: false,
				buttons: [{
					title: "Apply",
					autofocus: true,
					onclick: function () {
						var inputs = document.getElementById("fields").getElementsByTagName("input");
						var length = inputs.length;
						var payload = {
							"id": id,
							"panel": false
						}
						for (var i = 0; i < length; i++) {
							payload[inputs[i].name] = inputs[i].value;
						}
						post("/cmds/adjustHP", function (xhttp) {
							if (xhttp.status == 200) {
								document.getElementById("board-area").innerHTML = xhttp.responseText;
							}
							closeSimpleModal();
						}, JSON.stringify(payload));
					},
				}]
			});
		}
	}, JSON.stringify({
		"id": id,
		"panel": true
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
		buttons: [{
			title: "Delete",
			onclick: function () {
				post("/cmds/deleteCombatant", function (xhttp) {
					if (xhttp.status == 200) {
						document.getElementById("board-area").innerHTML = xhttp.responseText;
					}
					closeSimpleModal();
				}, JSON.stringify({
					"id": id
				}));
			},
		}]
	});
}

function addNote(id) {
	post("/cmds/addNote", function (xhttp) {
		if (xhttp.status == 200) {
			simpleModal({
				content: xhttp.responseText,
				wantAutoFocus: false,
				buttons: [{
					title: "Add",
					autofocus: true,
					onclick: function () {
						post("/cmds/addNote", function (xhttp) {
							if (xhttp.status == 200) {
								document.getElementById("board-area").innerHTML = xhttp.responseText;
							}
							closeSimpleModal();
						}, JSON.stringify(getNotePayload(id)));
					},
				}]
			});
		}
	}, JSON.stringify({
		"id": id,
		"panel": true
	}));
}

function getNotePayload(id) {
	var inputs = document.getElementById("fields").getElementsByTagName("input");
	var length = inputs.length;
	var payload = {
		"id": id,
		"panel": false
	}
	for (var i = 0; i < length; i++) {
		if (inputs[i].type == "checkbox") {
			payload[inputs[i].name] = inputs[i].checked;
		} else {
			payload[inputs[i].name] = inputs[i].value;
		}
	}
	inputs = document.getElementById("fields").getElementsByTagName("select");
	length = inputs.length;
	for (var i = 0; i < length; i++) {
		payload[inputs[i].name] = inputs[i].value;
	}
	return payload;
}

function editNote(id, index) {
	post("/cmds/editNote", function (xhttp) {
		if (xhttp.status == 200) {
			simpleModal({
				content: xhttp.responseText,
				wantAutoFocus: false,
				buttons: [{
					title: "Change",
					autofocus: true,
					onclick: function () {
						post("/cmds/editNote", function (xhttp) {
							if (xhttp.status == 200) {
								document.getElementById("board-area").innerHTML = xhttp.responseText;
							}
							closeSimpleModal();
						}, JSON.stringify(getNotePayload(id)));
					},
				}]
			});
		}
	}, JSON.stringify({
		"id": id,
		"index": index,
		"panel": true
	}));
}

function deleteNote(id, index) {
	post("/cmds/deleteNote", function (xhttp) {
		if (xhttp.status == 200) {
			document.getElementById("board-area").innerHTML = xhttp.responseText;
		}
	}, JSON.stringify({
		"id": id,
		"index": index
	}));
}

function nextTurn() {
	post("/cmds/nextTurn", function (xhttp) {
		if (xhttp.status == 200) {
			document.getElementById("board-area").innerHTML = xhttp.responseText;
			adjustRound(xhttp);
		}
	});
}

function sendSimpleCommand(cmd, id) {
	post("/cmds/" + cmd, function (xhttp) {
		if (xhttp.status == 200) {
			document.getElementById("board-area").innerHTML = xhttp.responseText;
		}
	}, JSON.stringify({
		"id": id
	}));
}

function simpleModal(options) {
	var dialog = document.getElementById("simple-modal-dialog");
	dialog.querySelector(".modal-content").innerHTML = options.content;
	var buttons = dialog.querySelector(".modal-buttons");
	buttons.innerHTML = "";
	var needAutofocus = true;
	var len = options.buttons.length;
	for (var i = 0; i < len; i++) {
		var button = document.createElement("button");
		button.onclick = options.buttons[i].onclick;
		if (options.buttons[i].autofocus) {
			needAutofocus = false;
			dialog.default_button = button
			if (options.wantAutoFocus) {
				button.autofocus = true;
			}
		}
		button.appendChild(document.createTextNode(options.buttons[i].title));
		buttons.appendChild(button);
	}
	var cancel = document.createElement("button");
	cancel.onclick = closeSimpleModal;
	if (needAutofocus) {
		dialog.default_button = cancel;
		if (options.wantAutoFocus) {
			cancel.autofocus = true;
		}
	}
	cancel.appendChild(document.createTextNode("Cancel"));
	dialog.cancel_button = cancel;
	buttons.appendChild(cancel);
	document.getElementById("modal-overlay").classList.remove("closed");
}

function closeSimpleModal() {
	document.getElementById("modal-overlay").classList.add("closed");
}

function isModalClosed() {
	return document.getElementById("modal-overlay").classList.contains("closed");
}

function isModalOpen() {
	return !isModalClosed();
}

function isActiveElementInput() {
	return document.activeElement !== undefined && document.activeElement.nodeName == "INPUT";
}

function isForGlobalKeyHandler() {
	return !isActiveElementInput() && isModalClosed();
}

function post(url, callback, data) {
	var xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function () {
		if (this.readyState == 4) {
			callback(xhttp);
		}
	};
	xhttp.open("POST", url, true);
	xhttp.setRequestHeader("Content-type", "application/json");
	xhttp.send(data);
}

function handleGlobalKeyDown(event) {
	if (isForGlobalKeyHandler()) {
		if (event.code == "KeyN") {
			event.stopPropagation();
			nextTurn();
		} else {
			handleUpDownArrows(event);
		}
	} else if (isModalClosed()) {
		handleUpDownArrows(event);
	}
}

function handleUpDownArrows(event) {
	switch (event.code) {
		case "ArrowDown":
			event.stopPropagation();
			event.preventDefault();
			showNextEntity();
			break;
		case "ArrowUp":
			event.stopPropagation();
			event.preventDefault();
			showPreviousEntity();
			break;
	}
}

function handleDefaultButton(event) {
	if (isModalOpen()) {
		switch (event.code) {
			case "Enter":
			case "NumpadEnter":
				var defButton = document.getElementById("simple-modal-dialog").default_button;
				if (defButton !== undefined) {
					event.stopPropagation();
					defButton.click();
				}
				break;
			case "Escape":
				var cancelButton = document.getElementById("simple-modal-dialog").cancel_button;
				if (cancelButton !== undefined) {
					event.stopPropagation();
					cancelButton.click();
				}
				break;
		}
	}
}

function adjustRound(xhttp) {
	var text;
	var round = xhttp.getResponseHeader("round");
	if (round < 1) {
		text = "Awaiting Initiative";
	} else {
		text = "Round " + round;
	}
	document.getElementById("turn").textContent = text;
	if (xhttp.getResponseHeader("new_round") == "true") {
		window.speechSynthesis.cancel();
		window.speechSynthesis.speak(new SpeechSynthesisUtterance("start of " + text));
	}
}

var combatantDNDType = "application/combatant";
var dragStartCombatant;
var dragStartIndex;
var dragCurrentOrder = [];
var dragCache;

function dragStartHandler(event) {
	dragCurrentOrder = [];
	dragStartCombatant = findCombatantID(event.target);
	if (dragStartCombatant != "") {
		event.dataTransfer.setData(combatantDNDType, dragStartCombatant);
		event.dataTransfer.effectAllowed = "move";
		var board = document.getElementById("board")
		var divs = board.getElementsByTagName("div");
		var length = divs.length;
		var last = "x";
		dragCache = new Map();
		for (var i = 0; i < length; i++) {
			var cid = divs[i].getAttribute("cid");
			if (cid != null) {
				if (cid != last) {
					last = cid;
					dragCurrentOrder.push(cid);
				}
			}
			dragCache.set(divs[i].className + cid, {
				offsetTop: divs[i].offsetTop,
				offsetHeight: divs[i].offsetHeight
			});
		}
		dragStartIndex = dragCurrentOrder.indexOf(dragStartCombatant);
	} else {
		event.preventDefault();
	}
}

function dragEnterHandler(event) {
	if (acceptableDrag(event)) {
		event.preventDefault();
		return;
	}
}

function dragOverHandler(event) {
	event.preventDefault();
	removePreviousDragMarkers();
	event.dataTransfer.dropEffect = "none";
	if (acceptableDrag(event)) {
		var i = dragToIndex(event);
		if (i != -1) {
			if (i == dragCurrentOrder.length) {
				mark(dragCurrentOrder[i - 1], "drag-marker-bottom");
			} else {
				mark(dragCurrentOrder[i], "drag-marker-top");
			}
			event.dataTransfer.dropEffect = "move";
		}
	}
}

function dragLeaveHandler(event) {
	removePreviousDragMarkers();
}

function dropHandler(event) {
	removePreviousDragMarkers();
	event.preventDefault();
	if (acceptableDrag(event)) {
		var i = dragToIndex(event);
		if (i != -1) {
			var pastEnd = i == dragCurrentOrder.length;
			var cid = dragCurrentOrder.splice(dragStartIndex, 1)[0];
			if (pastEnd) {
				dragCurrentOrder.push(cid);
			} else {
				if (dragStartIndex < i) {
					i--;
				}
				dragCurrentOrder.splice(i, 0, cid);
			}
			post("/cmds/reorder", function (xhttp) {
				if (xhttp.status == 200) {
					document.getElementById("board-area").innerHTML = xhttp.responseText;
				}
			}, JSON.stringify({
				"order": dragCurrentOrder.slice(0)
			}));
		}
	}
}

function dragEndHandler(event) {
	removePreviousDragMarkers();
	dragCurrentOrder = [];
}

function dragToIndex(event) {
	var node = findCombatantNode(event.toElement);
	if (node == null) {
		return -1;
	}
	let cid = node.getAttribute("cid");
	var i = dragCurrentOrder.indexOf(cid);
	let cache = dragCache.get(node.className + cid);
	if (event.clientY - (cache.offsetTop + cache.offsetHeight / 2) > 0) {
		i++;
	}
	if (i == dragStartIndex || i == dragStartIndex + 1) {
		return -1;
	}
	return i;
}

function mark(cid, className) {
	var board = document.getElementById("board")
	var divs = board.getElementsByTagName("div");
	var length = divs.length;
	for (var i = 0; i < length; i++) {
		if (divs[i].getAttribute("cid") == cid) {
			divs[i].classList.add(className);
		}
	}
}

function removePreviousDragMarkers() {
	var board = document.getElementById("board")
	removeClassFromChildren(board, "drag-marker-top");
	removeClassFromChildren(board, "drag-marker-bottom");
}

function removeClassFromChildren(parent, className) {
	var elems = parent.children;
	var length = elems.length;
	for (var i = 0; i < length; i++) {
		elems[i].classList.remove(className);
	}
}

function acceptableDrag(event) {
	var items = event.dataTransfer.items;
	for (var i = 0; i < items.length; i++) {
		var item = items[i];
		if (item.type == combatantDNDType) {
			var targetID = findCombatantID(event.toElement);
			return (targetID != "" && targetID != dragStartCombatant);
		}
	}
	return false;
}

function findCombatantID(elem) {
	var node = findCombatantNode(elem);
	return node == null ? "" : node.getAttribute("cid");
}

function findCombatantNode(elem) {
	while (elem !== undefined && elem != null && elem.getAttribute !== undefined && elem.getAttribute("cid") == null) {
		elem = elem.parentNode;
	}
	if (elem !== undefined && elem != null && elem.getAttribute !== undefined) {
		return elem;
	}
	return null;
}

function showPreviousEntity() {
	var elems = document.getElementById("library").children;
	var length = elems.length;
	var selectedElem;
	for (var i = 0; i < length; i++) {
		if (elems[i].classList.contains("library-selected")) {
			if (i != 0) {
				for (var j = i - 1; j >= 0; j--) {
					if (!elems[j].classList.contains("hide")) {
						showEntity(elems[j]);
						break;
					}
				}
			}
			break;
		}
	}
}

function showNextEntity() {
	var elems = document.getElementById("library").children;
	var length = elems.length;
	var selectedElem;
	for (var i = 0; i < length; i++) {
		if (elems[i].classList.contains("library-selected")) {
			if (i != length - 1) {
				for (var j = i + 1; j < length; j++) {
					if (!elems[j].classList.contains("hide")) {
						showEntity(elems[j]);
						break;
					}
				}
			}
			break;
		}
	}
}

function selectedEntityElem() {
	var elems = document.getElementById("library").children;
	var length = elems.length;
	for (var i = 0; i < length; i++) {
		if (elems[i].classList.contains("library-selected")) {
			return elems[i];
		}
	}
	return null;
}

function showEntity(target, force) {
	if ((force === undefined || !force) && target.classList.contains("library-selected")) {
		return;
	}
	var elems = document.getElementById("library").children;
	var length = elems.length;
	var selectedElem;
	for (var i = 0; i < length; i++) {
		if (elems[i].classList.contains("library-selected")) {
			selectedElem = elems[i];
			break;
		}
	}
	post("/cmds/showEntity", function (xhttp) {
		if (xhttp.status == 200) {
			selectedElem.classList.remove("library-selected");
			target.classList.add("library-selected");
			target.scrollIntoViewIfNeeded(false);
			var detail = document.getElementById("detail");
			detail.innerHTML = xhttp.responseText;
			detail.children[0].scrollIntoView();
		}
	}, JSON.stringify({
		"id": target.getAttribute("mid")
	}));
}

function findEntityByID(id) {
	var elems = document.getElementById("library").children;
	var length = elems.length;
	for (var i = 0; i < length; i++) {
		if (elems[i].getAttribute("mid") == id) {
			return elems[i];
		}
	}
	return null;
}

function libraryFilterChanged(value) {
	var exact = null;
	var contains = null;
	var containsIndex = 999999;
	var elems = document.getElementById("library").children;
	var length = elems.length;
	var count = 0;
	if (value == "") {
		for (var i = 0; i < length; i++) {
			elems[i].classList.remove("hide");
		}
		count = length;
	} else {
		value = value.toLowerCase();
		for (var i = 0; i < length; i++) {
			var e = elems[i];
			var name = e.getAttribute("name");
			var index = name.indexOf(value);
			if (index == -1) {
				if (!e.classList.contains("hide")) {
					e.classList.add("hide");
				}
			} else {
				count++;
				if (e.classList.contains("hide")) {
					e.classList.remove("hide");
				}
				if (exact === null) {
					if (name == value) {
						exact = e;
					} else if (index < containsIndex) {
						contains = e;
						containsIndex = index;
					}
				}
			}
		}
	}
	document.getElementById("filter-count").textContent = count;
	if (exact !== null) {
		showEntity(exact);
	} else if (contains !== null) {
		showEntity(contains);
	}
	scrollLibrarySelectionIntoViewIfNeeded();
}

function clearLibraryFilterOnEnter(event) {
	if (event.code == "Enter" || event.code == "NumpadEnter") {
		event.target.value = "";
		event.target.blur();
		libraryFilterChanged("");
	}
}

function scrollLibrarySelectionIntoViewIfNeeded() {
	var elems = document.getElementById("library").children;
	var length = elems.length;
	for (var i = 0; i < length; i++) {
		if (elems[i].classList.contains("library-selected")) {
			elems[i].scrollIntoViewIfNeeded(false);
			break;
		}
	}
}