/* This code goes through the babel output and 
   adds some color to make it easier for the reader */

function htmlEscape(rawString) {
    return rawString
        .replace(/&/g, "&amp;").replace(/"/g, "&quot;")
        .replace(/</g, "&lt;").replace(/>/g, "&gt;");
}

function prettifyAsciiOutput() {
    for (let pre of document.querySelectorAll("pre.example")) {
        let output = [];
        let lines = pre.textContent.split("\n");
        /* Find sections between markers */
        for (let i = 0, start = null; i < lines.length; i++) {
            if (start === null && /^_+$/.test(lines[i])) {
                start = i;
            } else if (start !== null && /^~+$/.test(lines[i])) {
                let html = htmlEscape(lines.slice(start+1, i).join("\n"));
                html = html.replace(/ &lt; /g, "<span>←</span>");
                html = html.replace(/ &gt; /g, "<span>→</span>");
                html = html.replace(/ v /g, "<span>↓</span>");
                html = html.replace(/ \^ /g, "<span>↑</span>");
                html = html.replace(/ [AZ] /g, "<i>$&</i>");
                html = html.replace(/( @ )+/g, "<b>$&</b>");
                html = html.replace(/#+/g, "<strong>$&</strong>");
                output.push(html);
                start = null;
            } else if (start === null) {
                output.push(htmlEscape(lines[i]));
            }
        }
        pre.innerHTML = output.join("\n");
    }
}

requestAnimationFrame(prettifyAsciiOutput);
