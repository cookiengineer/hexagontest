// Code highlighting for https://www.redblobgames.com/pathfinding/
// Copyright 2014 Red Blob Games <redblobgames@gmail.com>
// License: Apache v2.0 <http://www.apache.org/licenses/LICENSE-2.0.html>

// I don't want to syntax highlight; instead, I want to highlight
// specific words to match up with diagrams. This code surrounds given
// words X by <span class="X">X</span>

// Take plain text (from a <pre> section) and a set of words, and turn
// the text into html with those words marked. This code is for my own
// use and assumes that the words are \w+ with no spaces or
// punctuation etc. NOTE: although I don't have full <tag> exclusion,
// I avoid highlighting words that are followed by some punctuation
// marks that show up inside tags. This code is fragile.
export function highlight_words(text, words) {
    const wordBoundary = String.raw`\b`;
    const notTagNameOrAfterTag = String.raw`(?<!<)`;
    const wordsToMatch = `(${words.join("|")})`;
    const notAttribute = String.raw`(?![-'"=])`;
    try {
        const pattern = new RegExp(
        wordBoundary
            + notTagNameOrAfterTag
            + wordsToMatch
            + wordBoundary
                + notAttribute, 'g');
        return text.replace(pattern, "<span class='$&'>$&</span>");
    } catch (e) {
        // NOTE: Safari and Firefox ESR don't support lookbehind regexp
        console.log("WARNING: browser doesn't support https://caniuse.com/js-regexp-lookbehind, skipping code highlighting");
        return text;
    }
    // TODO: Words inside a <tag> should not match.
}

function test_highlight_words() {
    function T(a, b) {
        if (a != b) {
            console.log("ASSERT failed:");
            console.log("   OUTPUT = ", a);
            console.log("   EXPECT = ", b);
        }
    }
    T(highlight_words("begin end", ['begin', 'end']),
                   "<span class='begin'>begin</span> <span class='end'>end</span>");
    T(highlight_words("begin middle funcall() end", ['middle', 'funcall']),
                   "begin <span class='middle'>middle</span> <span class='funcall'>funcall</span>() end");
    T(highlight_words("begin eol\nbol end", ['eol', 'bol']),
                   "begin <span class='eol'>eol</span>\n<span class='bol'>bol</span> end");
    T(highlight_words("begin object.method end", ['object', 'method']),
                   "begin <span class='object'>object</span>.<span class='method'>method</span> end");
    T(highlight_words("begin <tag object='foo'> end", ['object', 'method', 'object']),
                   "begin <tag object='foo'> end");
    T(highlight_words("begin <tag foo='object'> end", ['object', 'method', 'foo']),
                   "begin <tag foo='object'> end");
}

test_highlight_words();
