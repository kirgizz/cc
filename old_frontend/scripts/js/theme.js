function setCookieT(cname, cvalue, exdays) {
    var d = new Date();
    d.setTime(d.getTime() + (exdays*24*60*60*1000));
    var expires = "expires="+ d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for(var i = 0; i <ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function changeTheme() {
    let body = document.body;
    if (body.classList.contains('dayT')) {
        body.classList.add('nightT');
        body.classList.remove('dayT');
        setCookieT("theme", "night", 360);
    } else {
        body.classList.add('dayT');
        body.classList.remove('nightT');
        setCookieT("theme", "day", 360);
    }
}

function onLoad() {
    if (getCookie("theme") === "night") {
        changeTheme();
    }

    let changes = document.getElementById('changes');
    let chist = APIGetChangesHistory();

    changes.innerHTML = '<pre>Last modifications: <br>';
    chist.forEach(function (item) {
        let time = Date.parse(item.At);
        let datech = new Date(time);
        changes.innerHTML = changes.innerHTML + '<b>' + item.User + '</b>' + ' at ' + datech.toLocaleString() + '<br>';
    });
    changes.innerHTML = changes.innerHTML + '</pre>';
}

function logout() {
    setCookieT("ssid", "", 360);
    window.location.replace("/login");
}
