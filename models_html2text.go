package mail

//"""html2text: Turn HTML into equivalent Markdown-structured text."""
//__version__ = "2.36"
//__author__ = "Aaron Swartz (me@aaronsw.com)"
//__copyright__ = "(C) 2004-2008 Aaron Swartz. GNU GPL 3."
//__contributors__ = ["Martin 'Joey' Schulze",
//                    "Ricardo Reyes", "Kevin Jay North"]
//if not hasattr(__builtins__, 'True'):
//    True, False = 1, 0
//import re
//import sys
//import urllib
//import htmlentitydefs
//import codecs
//import sgmllib
//import urlparse
//sgmllib.charref = re.compile('&#([xX]?[0-9a-fA-F]+)[^0-9a-fA-F]')
//try:
//    from textwrap import wrap
//except:
//    pass
//UNICODE_SNOB = 0
//LINKS_EACH_PARAGRAPH = 0
//BODY_WIDTH = 78
//SKIP_INTERNAL_LINKS = False
func Name2cp(k interface{}) {
	//    if k == 'apos':
	//        return ord("'")
	//    if hasattr(htmlentitydefs, "name2codepoint"):  # requires Python 2.3
	//        return htmlentitydefs.name2codepoint[k]
	//    else:
	//        k = htmlentitydefs.entitydefs[k]
	//        if k.startswith("&#") and k.endswith(";"):
	//            return int(k[2:-1])  # not in latin-1
	//        return ord(codecs.latin_1_decode(k)[0])
}

//unifiable = {'rsquo': "'", 'lsquo': "'", 'rdquo': '"', 'ldquo': '"',
//             'copy': '(C)', 'mdash': '--', 'nbsp': ' ', 'rarr': '->', 'larr': '<-', 'middot': '*',
//             'ndash': '-', 'oelig': 'oe', 'aelig': 'ae',
//             'agrave': 'a', 'aacute': 'a', 'acirc': 'a', 'atilde': 'a', 'auml': 'a', 'aring': 'a',
//             'egrave': 'e', 'eacute': 'e', 'ecirc': 'e', 'euml': 'e',
//             'igrave': 'i', 'iacute': 'i', 'icirc': 'i', 'iuml': 'i',
//             'ograve': 'o', 'oacute': 'o', 'ocirc': 'o', 'otilde': 'o', 'ouml': 'o',
//             'ugrave': 'u', 'uacute': 'u', 'ucirc': 'u', 'uuml': 'u'}
//unifiable_n = {}
//for k in unifiable.keys():
//    unifiable_n[name2cp(k)] = unifiable[k]
func Charref(name interface{}) {
	//    if name[0] in ['x', 'X']:
	//        c = int(name[1:], 16)
	//    else:
	//        c = int(name)
	//    if not UNICODE_SNOB and c in unifiable_n.keys():
	//        return unifiable_n[c]
	//    else:
	//        return unichr(c)
}
func Entityref(c interface{}) {
	//    if not UNICODE_SNOB and c in unifiable.keys():
	//        return unifiable[c]
	//    else:
	//        try:
	//            name2cp(c)
	//        except KeyError:
	//            return "&" + c
	//        else:
	//            return unichr(name2cp(c))
}
func ReplaceEntities(s interface{}) {
	//    s = s.group(1)
	//    if s[0] == "#":
	//        return charref(s[1:])
	//    else:
	//        return entityref(s)
}

//r_unescape = re.compile(r"&(#?[xX]?(?:[0-9a-fA-F]+|\w{1,8}));")
func Unescape(s interface{}) {
	//    return r_unescape.sub(replaceEntities, s)
}
func Fixattrs(attrs interface{}) {
	//    if not attrs:
	//        return attrs
	//    newattrs = []
	//    for attr in attrs:
	//        newattrs.append((attr[0], unescape(attr[1])))
	//    return newattrs
}
func Onlywhite(line interface{}) {
	//    """Return true if the line does only consist of whitespace characters."""
	//    for c in line:
	//        if c is not ' ' and c is not '  ':
	//            return c is ' '
	//    return line
}
func Optwrap(text interface{}) {
	//    """Wrap all paragraphs in the provided text."""
	//    if not BODY_WIDTH:
	//        return text
	//    assert wrap, "Requires Python 2.3."
	//    result = ''
	//    newlines = 0
	//    for para in text.split("\n"):
	//        if len(para) > 0:
	//            if para[0] is not ' ' and para[0] is not '-' and para[0] is not '*':
	//                for line in wrap(para, BODY_WIDTH):
	//                    result += line + "\n"
	//                result += "\n"
	//                newlines = 2
	//            else:
	//                if not onlywhite(para):
	//                    result += para + "\n"
	//                    newlines = 1
	//        else:
	//            if newlines < 2:
	//                result += "\n"
	//                newlines += 1
	//    return result
}
func Hn(tag interface{}) {
	//    if tag[0] == 'h' and len(tag) == 2:
	//        try:
	//            n = int(tag[1])
	//            if n in range(1, 10):
	//                return n
	//        except ValueError:
	//            return 0
}
func Init(self interface{}, out interface{}, baseurl interface{}) {
	//        sgmllib.SGMLParser.__init__(self)
	//        if out is None:
	//            self.out = self.outtextf
	//        else:
	//            self.out = out
	//        self.outtext = u''
	//        self.quiet = 0
	//        self.p_p = 0
	//        self.outcount = 0
	//        self.start = 1
	//        self.space = 0
	//        self.a = []
	//        self.astack = []
	//        self.acount = 0
	//        self.list = []
	//        self.blockquote = 0
	//        self.pre = 0
	//        self.startpre = 0
	//        self.lastWasNL = 0
	//        self.abbr_title = None  # current abbreviation definition
	//        self.abbr_data = None  # last inner HTML (for abbr being defined)
	//        self.abbr_list = {}  # stack of abbreviations to write later
	//        self.baseurl = baseurl
}
func Outtextf(self interface{}, s interface{}) {
	//        self.outtext += s
}
func Close(self interface{}) {
	//        sgmllib.SGMLParser.close(self)
	//        self.pbr()
	//        self.o('', 0, 'end')
	//        return self.outtext
}
func HandleCharref(self interface{}, c interface{}) {
	//        self.o(charref(c))
}
func HandleEntityref(self interface{}, c interface{}) {
	//        self.o(entityref(c))
}
func UnknownStarttag(self interface{}, tag interface{}, attrs interface{}) {
	//        self.handle_tag(tag, attrs, 1)
}
func UnknownEndtag(self interface{}, tag interface{}) {
	//        self.handle_tag(tag, None, 0)
}
func PreviousIndex(self interface{}, attrs interface{}) {
	//        """ returns the index of certain set of attributes (of a link) in the
	//            self.a list
	//
	//            If the set of attributes is not found, returns None
	//        """
	//        if not attrs.has_key('href'):
	//            return None
	//        i = -1
	//        for a in self.a:
	//            i += 1
	//            match = 0
	//
	//            if a.has_key('href') and a['href'] == attrs['href']:
	//                if a.has_key('title') or attrs.has_key('title'):
	//                    if (a.has_key('title') and attrs.has_key('title') and
	//                            a['title'] == attrs['title']):
	//                        match = True
	//                else:
	//                    match = True
	//
	//            if match:
	//                return i
}
func HandleTag(self interface{}, tag interface{}, attrs interface{}, start interface{}) {
	//        attrs = fixattrs(attrs)
	//        if hn(tag):
	//            self.p()
	//            if start:
	//                self.o(hn(tag)*"#" + ' ')
	//        if tag in ['p', 'div']:
	//            self.p()
	//        if tag == "br" and start:
	//            self.o("  \n")
	//        if tag == "hr" and start:
	//            self.p()
	//            self.o("* * *")
	//            self.p()
	//        if tag in ["head", "style", 'script']:
	//            if start:
	//                self.quiet += 1
	//            else:
	//                self.quiet -= 1
	//        if tag in ["body"]:
	//            self.quiet = 0  # sites like 9rules.com never close <head>
	//        if tag == "blockquote":
	//            if start:
	//                self.p()
	//                self.o('> ', 0, 1)
	//                self.start = 1
	//                self.blockquote += 1
	//            else:
	//                self.blockquote -= 1
	//                self.p()
	//        if tag in ['em', 'i', 'u']:
	//            self.o("_")
	//        if tag in ['strong', 'b']:
	//            self.o("**")
	//        if tag == "code" and not self.pre:
	//            self.o('`')  # TODO: `` `this` ``
	//        if tag == "abbr":
	//            if start:
	//                attrsD = {}
	//                for (x, y) in attrs:
	//                    attrsD[x] = y
	//                attrs = attrsD
	//
	//                self.abbr_title = None
	//                self.abbr_data = ''
	//                if attrs.has_key('title'):
	//                    self.abbr_title = attrs['title']
	//            else:
	//                if self.abbr_title != None:
	//                    self.abbr_list[self.abbr_data] = self.abbr_title
	//                    self.abbr_title = None
	//                self.abbr_data = ''
	//        if tag == "a":
	//            if start:
	//                attrsD = {}
	//                for (x, y) in attrs:
	//                    attrsD[x] = y
	//                attrs = attrsD
	//                if attrs.has_key('href') and not (SKIP_INTERNAL_LINKS and attrs['href'].startswith('#')):
	//                    self.astack.append(attrs)
	//                    self.o("[")
	//                else:
	//                    self.astack.append(None)
	//            else:
	//                if self.astack:
	//                    a = self.astack.pop()
	//                    if a:
	//                        i = self.previousIndex(a)
	//                        if i is not None:
	//                            a = self.a[i]
	//                        else:
	//                            self.acount += 1
	//                            a['count'] = self.acount
	//                            a['outcount'] = self.outcount
	//                            self.a.append(a)
	//                        self.o("][" + `a['count']` + "]")
	//        if tag == "img" and start:
	//            attrsD = {}
	//            for (x, y) in attrs:
	//                attrsD[x] = y
	//            attrs = attrsD
	//            if attrs.has_key('src'):
	//                attrs['href'] = attrs['src']
	//                alt = attrs.get('alt', '')
	//                i = self.previousIndex(attrs)
	//                if i is not None:
	//                    attrs = self.a[i]
	//                else:
	//                    self.acount += 1
	//                    attrs['count'] = self.acount
	//                    attrs['outcount'] = self.outcount
	//                    self.a.append(attrs)
	//                self.o("![")
	//                self.o(alt)
	//                self.o("]["+`attrs['count']`+"]")
	//        if tag == 'dl' and start:
	//            self.p()
	//        if tag == 'dt' and not start:
	//            self.pbr()
	//        if tag == 'dd' and start:
	//            self.o('    ')
	//        if tag == 'dd' and not start:
	//            self.pbr()
	//        if tag in ["ol", "ul"]:
	//            if start:
	//                self.list.append({'name': tag, 'num': 0})
	//            else:
	//                if self.list:
	//                    self.list.pop()
	//
	//            self.p()
	//        if tag == 'li':
	//            if start:
	//                self.pbr()
	//                if self.list:
	//                    li = self.list[-1]
	//                else:
	//                    li = {'name': 'ul', 'num': 0}
	//                # TODO: line up <ol><li>s > 9 correctly.
	//                self.o("  "*len(self.list))
	//                if li['name'] == "ul":
	//                    self.o("* ")
	//                elif li['name'] == "ol":
	//                    li['num'] += 1
	//                    self.o(`li['num']`+". ")
	//                self.start = 1
	//            else:
	//                self.pbr()
	//        if tag in ["table", "tr"] and start:
	//            self.p()
	//        if tag == 'td':
	//            self.pbr()
	//        if tag == "pre":
	//            if start:
	//                self.startpre = 1
	//                self.pre = 1
	//            else:
	//                self.pre = 0
	//            self.p()
}
func Pbr(self interface{}) {
	//        if self.p_p == 0:
	//            self.p_p = 1
}
func P(self interface{}) {
	//        self.p_p = 2
}
func O(self interface{}, data interface{}, puredata interface{}, force interface{}) {
	//        if self.abbr_data is not None:
	//            self.abbr_data += data
	//        if not self.quiet:
	//            if puredata and not self.pre:
	//                data = re.sub('\s+', ' ', data)
	//                if data and data[0] == ' ':
	//                    self.space = 1
	//                    data = data[1:]
	//            if not data and not force:
	//                return
	//
	//            if self.startpre:
	//                # self.out(" :") #TODO: not output when already one there
	//                self.startpre = 0
	//
	//            bq = (">" * self.blockquote)
	//            if not (force and data and data[0] == ">") and self.blockquote:
	//                bq += " "
	//
	//            if self.pre:
	//                bq += "    "
	//                data = data.replace("\n", "\n"+bq)
	//
	//            if self.start:
	//                self.space = 0
	//                self.p_p = 0
	//                self.start = 0
	//
	//            if force == 'end':
	//                # It's the end.
	//                self.p_p = 0
	//                self.out("\n")
	//                self.space = 0
	//
	//            if self.p_p:
	//                self.out(('\n'+bq)*self.p_p)
	//                self.space = 0
	//
	//            if self.space:
	//                if not self.lastWasNL:
	//                    self.out(' ')
	//                self.space = 0
	//
	//            if self.a and ((self.p_p == 2 and LINKS_EACH_PARAGRAPH) or force == "end"):
	//                if force == "end":
	//                    self.out("\n")
	//
	//                newa = []
	//                for link in self.a:
	//                    if self.outcount > link['outcount']:
	//                        self.out("   ["+`link['count']`+"]: " +
	//                                 urlparse.urljoin(self.baseurl, link['href']))
	//                        if link.has_key('title'):
	//                            self.out(" ("+link['title']+")")
	//                        self.out("\n")
	//                    else:
	//                        newa.append(link)
	//
	//                if self.a != newa:
	//                    # Don't need an extra line when nothing was done.
	//                    self.out("\n")
	//
	//                self.a = newa
	//
	//            if self.abbr_list and force == "end":
	//                for abbr, definition in self.abbr_list.items():
	//                    self.out("  *[" + abbr + "]: " + definition + "\n")
	//
	//            self.p_p = 0
	//            self.out(data)
	//            self.lastWasNL = data and data[-1] == '\n'
	//            self.outcount += 1
}
func HandleData(self interface{}, data interface{}) {
	//        if r'\/script>' in data:
	//            self.quiet -= 1
	//        self.o(data, 1)
}
func UnknownDecl(self interface{}, data interface{}) {
	//        pass
}
func Wrapwrite(text interface{}) {
	//def wrapwrite(text): sys.stdout.write(text.encode('utf8'))
}
func Html2textFile(html interface{}, out interface{}, baseurl interface{}) {
	//    h = _html2text(out, baseurl)
	//    h.feed(html)
	//    h.feed("")
	//    return h.close()
}
func Html2text(html interface{}, baseurl interface{}) {
	//    return optwrap(html2text_file(html, None, baseurl))
}

//if __name__ == "__main__":
//    baseurl = ''
//    if sys.argv[1:]:
//        arg = sys.argv[1]
//        if arg.startswith('http://'):
//            baseurl = arg
//            j = urllib.urlopen(baseurl)
//            try:
//                from feedparser import _getCharacterEncoding as enc
//            except ImportError:
//                def enc(x, y): return ('utf-8', 1)
//            text = j.read()
//            encoding = enc(j.headers, text)[0]
//            if encoding == 'us-ascii':
//                encoding = 'utf-8'
//            data = text.decode(encoding)
//
//        else:
//            encoding = 'utf8'
//            if len(sys.argv) > 2:
//                encoding = sys.argv[2]
//            f = open(arg, 'r')
//            try:
//                data = f.read().decode(encoding)
//            finally:
//                f.close()
//    else:
//        data = sys.stdin.read().decode('utf8')
//    wrapwrite(html2text(data, baseurl))
