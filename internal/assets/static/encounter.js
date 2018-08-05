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
						let inputs = document.getElementById("fields").getElementsByTagName("input");
						let length = inputs.length;
						let inits = []
						for (let i = 0; i < length; i++) {
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
						let payload = {
							"panel": false
						};
						let inputs = document.getElementById("fields").getElementsByTagName("input");
						let length = inputs.length;
						for (let i = 0; i < length; i++) {
							payload[inputs[i].name] = inputs[i].value;
						}
						inputs = document.getElementById("fields").getElementsByTagName("select");
						length = inputs.length;
						for (let i = 0; i < length; i++) {
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
	let payload = {
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
							let inputs = document.getElementById("fields").getElementsByTagName("input");
							let length = inputs.length;
							payload.panel = false;
							for (let i = 0; i < length; i++) {
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
						let inputs = document.getElementById("fields").getElementsByTagName("input");
						let length = inputs.length;
						let payload = {
							"id": id,
							"panel": false
						}
						for (let i = 0; i < length; i++) {
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
						let inputs = document.getElementById("fields").getElementsByTagName("input");
						let length = inputs.length;
						let payload = {
							"id": id,
							"panel": false
						}
						for (let i = 0; i < length; i++) {
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

function directAdjustHP(id, amount) {
	post("/cmds/adjustHP", function (xhttp) {
		if (xhttp.status == 200) {
			document.getElementById("board-area").innerHTML = xhttp.responseText;
		}
	}, JSON.stringify({
		"id": id,
		"panel": false,
		"adjust": amount
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
	let inputs = document.getElementById("fields").getElementsByTagName("input");
	let length = inputs.length;
	let payload = {
		"id": id,
		"panel": false
	}
	for (let i = 0; i < length; i++) {
		if (inputs[i].type == "checkbox") {
			payload[inputs[i].name] = inputs[i].checked;
		} else {
			payload[inputs[i].name] = inputs[i].value;
		}
	}
	inputs = document.getElementById("fields").getElementsByTagName("select");
	length = inputs.length;
	for (let i = 0; i < length; i++) {
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
	let dialog = document.getElementById("simple-modal-dialog");
	dialog.querySelector(".modal-content").innerHTML = options.content;
	let buttons = dialog.querySelector(".modal-buttons");
	buttons.innerHTML = "";
	let needAutofocus = true;
	let len = options.buttons.length;
	for (let i = 0; i < len; i++) {
		let button = document.createElement("button");
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
	let cancel = document.createElement("button");
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
	// Ensure the focus is not within the closed dialog.
	let active = document.activeElement;
	let cur = active;
	while (cur !== undefined && cur != null) {
		if (cur.classList.contains("closed")) {
			active.blur();
			break;
		}
		cur = cur.parentElement;
	}
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
	let xhttp = new XMLHttpRequest();
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
				let defButton = document.getElementById("simple-modal-dialog").default_button;
				if (defButton !== undefined) {
					event.stopPropagation();
					defButton.click();
				}
				break;
			case "Escape":
				let cancelButton = document.getElementById("simple-modal-dialog").cancel_button;
				if (cancelButton !== undefined) {
					event.stopPropagation();
					cancelButton.click();
				}
				break;
		}
	}
}

function adjustRound(xhttp) {
	let text;
	let round = xhttp.getResponseHeader("round");
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

let combatantDNDType = "application/combatant";
let dragStartCombatant;
let dragStartIndex;
let dragCurrentOrder = [];
let dragCache;

function dragStartHandler(event) {
	dragCurrentOrder = [];
	dragStartCombatant = findCombatantID(event.target);
	if (dragStartCombatant != "") {
		event.dataTransfer.setData(combatantDNDType, dragStartCombatant);
		event.dataTransfer.effectAllowed = "move";
		let board = document.getElementById("board")
		let divs = board.getElementsByTagName("div");
		let length = divs.length;
		let last = "x";
		dragCache = new Map();
		for (let i = 0; i < length; i++) {
			let cid = divs[i].getAttribute("cid");
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
		let i = dragToIndex(event);
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
		let i = dragToIndex(event);
		if (i != -1) {
			let pastEnd = i == dragCurrentOrder.length;
			let cid = dragCurrentOrder.splice(dragStartIndex, 1)[0];
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
	let node = findCombatantNode(event.toElement);
	if (node == null) {
		return -1;
	}
	let cid = node.getAttribute("cid");
	let i = dragCurrentOrder.indexOf(cid);
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
	let board = document.getElementById("board")
	let divs = board.getElementsByTagName("div");
	let length = divs.length;
	for (let i = 0; i < length; i++) {
		if (divs[i].getAttribute("cid") == cid) {
			divs[i].classList.add(className);
		}
	}
}

function removePreviousDragMarkers() {
	let board = document.getElementById("board")
	removeClassFromChildren(board, "drag-marker-top");
	removeClassFromChildren(board, "drag-marker-bottom");
}

function removeClassFromChildren(parent, className) {
	let elems = parent.children;
	let length = elems.length;
	for (let i = 0; i < length; i++) {
		elems[i].classList.remove(className);
	}
}

function acceptableDrag(event) {
	let items = event.dataTransfer.items;
	for (let i = 0; i < items.length; i++) {
		let item = items[i];
		if (item.type == combatantDNDType) {
			let targetID = findCombatantID(event.toElement);
			return (targetID != "" && targetID != dragStartCombatant);
		}
	}
	return false;
}

function findCombatantID(elem) {
	let node = findCombatantNode(elem);
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
	let elems = document.getElementById("library").children;
	let length = elems.length;
	let selectedElem;
	for (let i = 0; i < length; i++) {
		if (elems[i].classList.contains("library-selected")) {
			if (i != 0) {
				for (let j = i - 1; j >= 0; j--) {
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
	let elems = document.getElementById("library").children;
	let length = elems.length;
	let selectedElem;
	for (let i = 0; i < length; i++) {
		if (elems[i].classList.contains("library-selected")) {
			if (i != length - 1) {
				for (let j = i + 1; j < length; j++) {
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
	let elems = document.getElementById("library").children;
	let length = elems.length;
	for (let i = 0; i < length; i++) {
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
	let elems = document.getElementById("library").children;
	let length = elems.length;
	let selectedElem;
	for (let i = 0; i < length; i++) {
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
			let detail = document.getElementById("detail");
			detail.innerHTML = xhttp.responseText;
			detail.children[0].scrollIntoView();
		}
	}, JSON.stringify({
		"id": target.getAttribute("mid")
	}));
}

function findEntityByID(id) {
	let elems = document.getElementById("library").children;
	let length = elems.length;
	for (let i = 0; i < length; i++) {
		if (elems[i].getAttribute("mid") == id) {
			return elems[i];
		}
	}
	return null;
}

function libraryFilterChanged(value) {
	let exact = null;
	let contains = null;
	let containsIndex = 999999;
	let elems = document.getElementById("library").children;
	let length = elems.length;
	let count = 0;
	if (value == "") {
		for (let i = 0; i < length; i++) {
			elems[i].classList.remove("hide");
		}
		count = length;
	} else {
		value = value.toLowerCase();
		for (let i = 0; i < length; i++) {
			let e = elems[i];
			let name = e.getAttribute("name");
			let index = name.indexOf(value);
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
	let elems = document.getElementById("library").children;
	let length = elems.length;
	for (let i = 0; i < length; i++) {
		if (elems[i].classList.contains("library-selected")) {
			elems[i].scrollIntoViewIfNeeded(false);
			break;
		}
	}
}