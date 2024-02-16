// Copyright 2023 Specter Ops, Inc.
// 
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// 
// SPDX-License-Identifier: Apache-2.0

// Code generated by gen.go; DO NOT EDIT.
// This file was generated at 2023-04-18 20:46:35.896682 -0400 EDT m=+2.098209075
// using data from https://www.iana.org/assignments/media-types/text.csv
package mediatypes

const (
	Text1dInterleavedParityfec                     MediaType = "text/1d-interleaved-parityfec"         // [RFC6015],
	TextCacheManifest                              MediaType = "text/cache-manifest"                   // [W3C][Robin_Berjon],
	TextCalendar                                   MediaType = "text/calendar"                         // [RFC5545],
	TextCql                                        MediaType = "text/cql"                              // [HL7][Bryn_Rhodes],
	TextCqlExpression                              MediaType = "text/cql-expression"                   // [HL7][Bryn_Rhodes],
	TextCqlIdentifier                              MediaType = "text/cql-identifier"                   // [HL7][Bryn_Rhodes],
	TextCss                                        MediaType = "text/css"                              // [RFC2318],
	TextCsv                                        MediaType = "text/csv"                              // [RFC4180][RFC7111],
	TextCsvSchema                                  MediaType = "text/csv-schema"                       // [National_Archives_UK][David_Underdown],
	TextDirectoryDEPRECATEDByRFC6350               MediaType = "text/directory"                        // [RFC2425][RFC6350],
	TextDns                                        MediaType = "text/dns"                              // [RFC4027],
	TextEcmascriptOBSOLETEDInFavorOfTextJavascript MediaType = "text/ecmascript"                       // [RFC9239],
	TextEncaprtp                                   MediaType = "text/encaprtp"                         // [RFC6849],
	TextExample                                    MediaType = "text/example"                          // [RFC4735],
	TextFhirpath                                   MediaType = "text/fhirpath"                         // [HL7][Bryn_Rhodes],
	TextFlexfec                                    MediaType = "text/flexfec"                          // [RFC8627],
	TextFwdred                                     MediaType = "text/fwdred"                           // [RFC6354],
	TextGff3                                       MediaType = "text/gff3"                             // [Sequence_Ontology],
	TextGrammarRefList                             MediaType = "text/grammar-ref-list"                 // [RFC6787],
	TextHl7v2                                      MediaType = "text/hl7v2"                            // [HL7][Marc_Duteau],
	TextHtml                                       MediaType = "text/html"                             // [W3C][Robin_Berjon],
	TextJavascript                                 MediaType = "text/javascript"                       // [RFC9239],
	TextJcrCnd                                     MediaType = "text/jcr-cnd"                          // [Peeter_Piegaze],
	TextMarkdown                                   MediaType = "text/markdown"                         // [RFC7763],
	TextMizar                                      MediaType = "text/mizar"                            // [Jesse_Alama],
	TextN3                                         MediaType = "text/n3"                               // [W3C][Eric_Prudhommeaux],
	TextParameters                                 MediaType = "text/parameters"                       // [RFC7826],
	TextParityfec                                  MediaType = "text/parityfec"                        // [RFC3009],
	TextProvenanceNotation                         MediaType = "text/provenance-notation"              // [W3C][Ivan_Herman],
	TextPrsFallensteinRst                          MediaType = "text/prs.fallenstein.rst"              // [Benja_Fallenstein],
	TextPrsLinesTag                                MediaType = "text/prs.lines.tag"                    // [John_Lines],
	TextPrsPropLogic                               MediaType = "text/prs.prop.logic"                   // [Hans-Dieter_A._Hiep],
	TextRaptorfec                                  MediaType = "text/raptorfec"                        // [RFC6682],
	TextRED                                        MediaType = "text/RED"                              // [RFC4102],
	TextRfc822Headers                              MediaType = "text/rfc822-headers"                   // [RFC6522],
	TextRtf                                        MediaType = "text/rtf"                              // [Paul_Lindner],
	TextRtpEncAescm128                             MediaType = "text/rtp-enc-aescm128"                 // [_3GPP],
	TextRtploopback                                MediaType = "text/rtploopback"                      // [RFC6849],
	TextRtx                                        MediaType = "text/rtx"                              // [RFC4588],
	TextSGML                                       MediaType = "text/SGML"                             // [RFC1874],
	TextShaclc                                     MediaType = "text/shaclc"                           // [W3C_SHACL_Community_Group][Vladimir_Alexiev],
	TextShex                                       MediaType = "text/shex"                             // [W3C][Eric_Prudhommeaux],
	TextSpdx                                       MediaType = "text/spdx"                             // [Linux_Foundation][Rose_Judge],
	TextStrings                                    MediaType = "text/strings"                          // [IEEE-ISTO-PWG-PPP],
	TextT140                                       MediaType = "text/t140"                             // [RFC4103],
	TextTabSeparatedValues                         MediaType = "text/tab-separated-values"             // [Paul_Lindner],
	TextTroff                                      MediaType = "text/troff"                            // [RFC4263],
	TextTurtle                                     MediaType = "text/turtle"                           // [W3C][Eric_Prudhommeaux],
	TextUlpfec                                     MediaType = "text/ulpfec"                           // [RFC5109],
	TextUriList                                    MediaType = "text/uri-list"                         // [RFC2483],
	TextVcard                                      MediaType = "text/vcard"                            // [RFC6350],
	TextVndA                                       MediaType = "text/vnd.a"                            // [Regis_Dehoux],
	TextVndAbc                                     MediaType = "text/vnd.abc"                          // [Steve_Allen],
	TextVndAsciiArt                                MediaType = "text/vnd.ascii-art"                    // [Kim_Scarborough],
	TextVndCurl                                    MediaType = "text/vnd.curl"                         // [Robert_Byrnes],
	TextVndDebianCopyright                         MediaType = "text/vnd.debian.copyright"             // [Charles_Plessy],
	TextVndDMClientScript                          MediaType = "text/vnd.DMClientScript"               // [Dan_Bradley],
	TextVndDvbSubtitle                             MediaType = "text/vnd.dvb.subtitle"                 // [Peter_Siebert][Michael_Lagally],
	TextVndEsmertecThemeDescriptor                 MediaType = "text/vnd.esmertec.theme-descriptor"    // [Stefan_Eilemann],
	TextVndExchangeable                            MediaType = "text/vnd.exchangeable"                 // [Martin_Cizek],
	TextVndFamilysearchGedcom                      MediaType = "text/vnd.familysearch.gedcom"          // [Gordon_Clarke],
	TextVndFiclabFlt                               MediaType = "text/vnd.ficlab.flt"                   // [Steve_Gilberd],
	TextVndFly                                     MediaType = "text/vnd.fly"                          // [John-Mark_Gurney],
	TextVndFmiFlexstor                             MediaType = "text/vnd.fmi.flexstor"                 // [Kari_E._Hurtta],
	TextVndGml                                     MediaType = "text/vnd.gml"                          // [Mi_Tar],
	TextVndGraphviz                                MediaType = "text/vnd.graphviz"                     // [John_Ellson],
	TextVndHans                                    MediaType = "text/vnd.hans"                         // [Hill_Hanxv],
	TextVndHgl                                     MediaType = "text/vnd.hgl"                          // [Heungsub_Lee],
	TextVndIn3d3dml                                MediaType = "text/vnd.in3d.3dml"                    // [Michael_Powers],
	TextVndIn3dSpot                                MediaType = "text/vnd.in3d.spot"                    // [Michael_Powers],
	TextVndIPTCNewsML                              MediaType = "text/vnd.IPTC.NewsML"                  // [IPTC],
	TextVndIPTCNITF                                MediaType = "text/vnd.IPTC.NITF"                    // [IPTC],
	TextVndLatexZ                                  MediaType = "text/vnd.latex-z"                      // [Mikusiak_Lubos],
	TextVndMotorolaReflex                          MediaType = "text/vnd.motorola.reflex"              // [Mark_Patton],
	TextVndMsMediapackage                          MediaType = "text/vnd.ms-mediapackage"              // [Jan_Nelson],
	TextVndNet2phoneCommcenterCommand              MediaType = "text/vnd.net2phone.commcenter.command" // [Feiyu_Xie],
	TextVndRadisysMsmlBasicLayout                  MediaType = "text/vnd.radisys.msml-basic-layout"    // [RFC5707],
	TextVndSenxWarpscript                          MediaType = "text/vnd.senx.warpscript"              // [Pierre_Papin],
	TextVndSiUricatalogueOBSOLETEDByRequest        MediaType = "text/vnd.si.uricatalogue"              // [Nicholas_Parks_Young],
	TextVndSunJ2meAppDescriptor                    MediaType = "text/vnd.sun.j2me.app-descriptor"      // [Gary_Adams],
	TextVndSosi                                    MediaType = "text/vnd.sosi"                         // [Petter_Reinholdtsen],
	TextVndTrolltechLinguist                       MediaType = "text/vnd.trolltech.linguist"           // [David_Lee_Lambert],
	TextVndWapSi                                   MediaType = "text/vnd.wap.si"                       // [WAP-Forum],
	TextVndWapSl                                   MediaType = "text/vnd.wap.sl"                       // [WAP-Forum],
	TextVndWapWml                                  MediaType = "text/vnd.wap.wml"                      // [Peter_Stark],
	TextVndWapWmlscript                            MediaType = "text/vnd.wap.wmlscript"                // [Peter_Stark],
	TextVtt                                        MediaType = "text/vtt"                              // [W3C][Silvia_Pfeiffer],
	TextWgsl                                       MediaType = "text/wgsl"                             // [W3C][David_Neto],
	TextXml                                        MediaType = "text/xml"                              // [RFC7303],
	TextXmlExternalParsedEntity                    MediaType = "text/xml-external-parsed-entity"       // [RFC7303],
)
