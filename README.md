# ccc (change-character-code)
It is a library for changing the character code of the string.

## [Support Character Code](https://encoding.spec.whatwg.org/#names-and-labels)
### The Encoding
| Character Code | Supported | Sample |
| --- | :---: | --- |
| UTF-8 |  |  |

### Legacy single-byte encodings
| Character Code | Supported | Sample |
| --- | :---: | --- |
| IBM866 |  |  |
| ISO-8859-2 |  |  |
| ISO-8859-3 |  |  |
| ISO-8859-4 |  |  |
| ISO-8859-5 |  |  |
| ISO-8859-6 |  |  |
| ISO-8859-7 |  |  |
| ISO-8859-8 |  |  |
| ISO-8859-8-I |  |  |
| ISO-8859-10 |  |  |
| ISO-8859-13 |  |  |
| ISO-8859-14 |  |  |
| ISO-8859-15 |  |  |
| ISO-8859-16 |  |  |
| KOI8-R |  |  |
| KOI8-U |  |  |
| macintosh |  |  |
| windows-874 |  |  |
| windows-1250 |  |  |
| windows-1251 |  |  |
| windows-1252 |  |  |
| windows-1253 |  |  |
| windows-1254 |  |  |
| windows-1255 |  |  |
| windows-1256 |  |  |
| windows-1257 |  |  |
| windows-1258 |  |  |
| x-mac-cyrillic |  |  |

### Legacy multi-byte Chinese (simplified) encodings
| Character Code | Supported | Sample |
| --- | :---: | --- |
| GBK | ✅ | ccc.Encode(string).GBK() |
| gb18030 | ✅ | ccc.Encode(string).GB18030() |

### Legacy multi-byte Chinese (traditional) encodings
| Character Code | Supported | Sample |
| --- | :---: | --- |
| Big5 | ✅ | ccc.Encode(string).Big5() |

### Legacy multi-byte Japanese encodings
| Character Code | Supported | Sample |
| --- | :---: | --- |
| EUC-JP | ✅ | ccc.Encode(string).EUC_JP() |
| ISO-2022-JP | ✅ | ccc.Encode(string).ISO_2022_JP() |
| Shift_JIS | ✅ | ccc.Encode(string).ShiftJIS() |

### Legacy multi-byte Korean encodings
| Character Code | Supported | Sample |
| --- | :---: | --- |
| EUC-KR |  |  |

### Legacy miscellaneous encodings
| Character Code | Supported | Sample |
| --- | :---: | --- |
| replacement |  |  |
| UTF-16BE |  |  |
| UTF-16LE |  |  |
| x-user-defined |  |  |

## License

MIT

## Author

Chihiro Suzurikawa (chihiros)
