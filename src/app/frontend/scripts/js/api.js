function APIGetDuty() {
    let xhr = new XMLHttpRequest();
    xhr.open('GET', 'api/getduties', false);
    xhr.send();

    if (xhr.readyState === 4 && xhr.status === 200) {
        return JSON.parse(xhr.responseText);
    }
    return JSON.parse('{}');
}

function APIGetChangesHistory() {
    let xhr = new XMLHttpRequest();
    xhr.open('GET', 'api/getchangeshistory', false);
    xhr.send();

    if (xhr.readyState === 4 && xhr.status === 200) {
        return JSON.parse(xhr.responseText);
    }
    return JSON.parse('{}');
}

function APIMakeDefaultSchedule(count, startFrom) {
    let xhr = new XMLHttpRequest();
    xhr.open('POST', 'api/makedefaultschedule', false);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.send(JSON.stringify({count: count, start_from: startFrom}));

    if (xhr.readyState === 4 && xhr.status === 200) {
        return JSON.parse(xhr.responseText);
    }
    return JSON.parse('[]');
}

function APIGetSchedule() {
    let xhr = new XMLHttpRequest();
    xhr.open('POST', 'api/getschedule', false);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.send(JSON.stringify({month: month + 1, year: year}));

    if (xhr.readyState === 4 && xhr.status === 200) {
        return JSON.parse(xhr.responseText);
        // alert('ok');
    }
    return JSON.parse('{}');
}

