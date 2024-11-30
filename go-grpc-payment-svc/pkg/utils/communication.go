package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/resendlabs/resend-go"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func SendSMS(to string, message string) error {
	link := "https://api.africastalking.com/version1/messaging"
	method := "POST"

	data := url.Values{}
	data.Set("username", "enyumba")
	data.Set("to", to)
	data.Set("message", message)
	body := strings.NewReader(data.Encode())
	client := &http.Client{}

	request, error := http.NewRequest(method, link, body)
	if error != nil {
		return error
	}
	live := "13f2ecd5c12eb5a403d51059a76f85dfee15ae2c40cc93d39305c8aac6546eaf"
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("apiKey", live)

	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err != nil {
		return err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	return nil
}

func SendEmail(receiverEmail string, subject string, title string, message string) error {
	apiKey := "re_2iCbBL9i_6V7BNKxHr7rZqrP974QtMGqX"

	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From: "Enyumba App <info@enyumba.com>",
		To:   []string{receiverEmail},
		Html: `
		<!DOCTYPE html
		PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
		<html xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml"
		xmlns:o="urn:schemas-microsoft-com:office:office" lang="en">

		<head>
		<meta name=x-apple-disable-message-reformatting>
		<meta http-equiv=X-UA-Compatible>
		<meta charset=utf-8>
		<meta name=viewport content=target-densitydpi=device-dpi>
		<meta content=true name=HandheldFriendly>
		<meta content=width=device-width name=viewport>
		<style type="text/css">
			table {
				border-collapse: separate;
				table-layout: fixed;
				mso-table-lspace: 0pt;
				mso-table-rspace: 0pt
			}

			table td {
				border-collapse: collapse
			}

			.ExternalClass {
				width: 100%
			}

			.ExternalClass,
			.ExternalClass p,
			.ExternalClass span,
			.ExternalClass font,
			.ExternalClass td,
			.ExternalClass div {
				line-height: 100%
			}

			* {
				line-height: inherit;
				text-size-adjust: 100%;
				-ms-text-size-adjust: 100%;
				-moz-text-size-adjust: 100%;
				-o-text-size-adjust: 100%;
				-webkit-text-size-adjust: 100%;
				-webkit-font-smoothing: antialiased;
				-moz-osx-font-smoothing: grayscale
			}

			html {
				-webkit-text-size-adjust: none !important
			}

			#innerTable img+div {
				display: none;
				display: none !important
			}

			img {
				Margin: 0;
				padding: 0;
				-ms-interpolation-mode: bicubic
			}

			h1,
			h2,
			h3,
			p,
			a {
				line-height: 1;
				overflow-wrap: normal;
				white-space: normal;
				word-break: break-word
			}

			a {
				text-decoration: none
			}

			h1,
			h2,
			h3,
			p {
				min-width: 100% !important;
				width: 100% !important;
				max-width: 100% !important;
				display: inline-block !important;
				border: 0;
				padding: 0;
				margin: 0
			}

			a[x-apple-data-detectors] {
				color: inherit !important;
				text-decoration: none !important;
				font-size: inherit !important;
				font-family: inherit !important;
				font-weight: inherit !important;
				line-height: inherit !important
			}

			a[href^="mailto"],
			a[href^="tel"],
			a[href^="sms"] {
				color: inherit;
				text-decoration: none
			}

			@media (min-width: 481px) {
				.hd {
					display: none !important
				}
			}

			@media (max-width: 480px) {
				.hm {
					display: none !important
				}
			}

			[style*="Albert Sans"] {
				font-family: 'Albert Sans', BlinkMacSystemFont, Segoe UI, Helvetica Neue, Arial, sans-serif !important;
			}

			@media only screen and (min-width: 481px) {
				.t3 {
					mso-line-height-alt: 45px !important;
					line-height: 45px !important;
					display: block !important
				}

				.t9 {
					padding-left: 50px !important;
					padding-bottom: 60px !important;
					padding-right: 50px !important
				}

				.t11 {
					padding-left: 50px !important;
					padding-bottom: 60px !important;
					padding-right: 50px !important;
					width: 500px !important
				}

				.t15,
				.t20 {
					padding-bottom: 25px !important
				}

				.t21 {
					line-height: 41px !important;
					font-size: 39px !important;
					letter-spacing: -1.56px !important
				}

				.t28 {
					padding: 48px 50px !important
				}

				.t30 {
					padding: 48px 50px !important;
					width: 500px !important
				}

				.t44,
				.t49 {
					padding-bottom: 44px !important
				}

				.t127,
				.t132 {
					padding-bottom: 45px !important
				}

				.t147 {
					padding-bottom: 60px !important;
					width: 130px !important
				}

				.t152 {
					padding-bottom: 60px !important
				}
			}
		</style>
		<!--[if !mso]><!-->
		<link href="https://fonts.googleapis.com/css2?family=Albert+Sans:wght@500;700;800&display=swap" rel="stylesheet"
			type="text/css">
		<!--<![endif]-->
		<!--[if mso]>
			<style type="text/css">
			div.t3{mso-line-height-alt:45px !important;line-height:45px !important;display:block !important}td.t11,td.t9{padding-left:50px !important;padding-bottom:60px !important;padding-right:50px !important}td.t15,td.t20{padding-bottom:25px !important}h1.t21{line-height:41px !important;font-size:39px !important;letter-spacing:-1.56px !important}td.t28,td.t30{padding:48px 50px !important}td.t44,td.t49{padding-bottom:44px !important}td.t127,td.t132{padding-bottom:45px !important}td.t147{padding-bottom:60px !important;width:130px !important}td.t152{padding-bottom:60px !important}
			</style>
			<![endif]-->
		<!--[if mso]>
			<xml>
			<o:OfficeDocumentSettings>
			<o:AllowPNG/>
			<o:PixelsPerInch>96</o:PixelsPerInch>
			</o:OfficeDocumentSettings>
			</xml>
			<![endif]-->
		</head>

		<body class=t0 style="min-width:100%;Margin:0px;padding:0px;background-color:#242424;">
		<div class=t1 style="background-color:#242424;">
			<table role=presentation width=100% cellpadding=0 cellspacing=0 border=0 align=center>
				<tr>
					<td class=t154 style="font-size:0;line-height:0;mso-line-height-rule:exactly;" valign=top align=center>
						<!--[if mso]>
			<v:background xmlns:v="urn:schemas-microsoft-com:vml" fill="true" stroke="false">
			<v:fill color=#242424 />
			</v:background>
			<![endif]-->
						<table role=presentation width=100% cellpadding=0 cellspacing=0 border=0 align=center id=innerTable>
							<tr>
								<td>
									<div class=t3 style="mso-line-height-rule:exactly;font-size:1px;display:none;">&nbsp;
									</div>
								</td>
							</tr>
							<tr>
								<td>
									<table class=t10 role=presentation cellpadding=0 cellspacing=0 align=center>
										<tr>
											<!--[if !mso]><!-->
											<td class=t11
												style="background-color:#F8F8F8;overflow:hidden;width:540px;padding:0 30px 40px 30px;">
												<!--<![endif]-->
												<!--[if mso]><td class=t11 style="background-color:#F8F8F8;overflow:hidden;width:600px;padding:0 30px 40px 30px;"><![endif]-->
												<table role=presentation width=100% cellpadding=0 cellspacing=0>
													<tr>
														<td>
															<table class=t146 role=presentation cellpadding=0 cellspacing=0
																align=left>
																<tr>
																	<!--[if !mso]><!-->
																	<td class=t147 style="width:80px;padding:40px 0 50px 0;">
																		<!--<![endif]-->
																		<!--[if mso]><td class=t147 style="width:80px;padding:0 0 50px 0;"><![endif]-->
																		<div style="font-size:0px;"><img class=t153
																				style="display:block;border:0;height:auto;width:150%;margin-top:30;max-width:150%;"

																				src=https://firebasestorage.googleapis.com/v0/b/enyumba-web.appspot.com/o/enyumba_word_logo.png?alt=media&token=0cbc2b34-f986-4936-b414-bf317c8088e7 />
																		</div>
																	</td>
																</tr>
															</table>
														</td>
													</tr>
													<tr>
														<td>
															<table class=t14 role=presentation cellpadding=0 cellspacing=0
																align=center>
																<tr>
																	<!--[if !mso]><!-->
																	<td class=t15 style="width:600px;padding:0 0 20px 0;">
																		<!--<![endif]-->
																		<!--[if mso]><td class=t15 style="width:600px;padding:0 0 20px 0;"><![endif]-->
																		<h1 class=t21
																			style="margin-bottom:0;Margin-bottom:0;font-family:BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif,'Albert Sans';line-height:28px;font-weight:800;font-style:normal;font-size:26px;text-decoration:none;text-transform:none;letter-spacing:-1.04px;direction:ltr;color:#191919;text-align:left;mso-line-height-rule:exactly;mso-text-raise:1px;">
																			` + title + `</h1>
																	</td>
																</tr>
															</table>
														</td>
													</tr>
													<tr>
														<td>
															<table class=t136 role=presentation cellpadding=0 cellspacing=0
																align=center>
																<tr>
																	<!--[if !mso]><!-->
																	<td class=t137 style="width:600px;padding:0 0 22px 0;">
																		<!--<![endif]-->
																		<!--[if mso]><td class=t137 style="width:600px;padding:0 0 22px 0;"><![endif]-->
																		<p class=t143
																			style="margin-bottom:0;Margin-bottom:0;font-family:BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif,'Albert Sans';line-height:22px;font-weight:500;font-style:normal;font-size:14px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#333333;text-align:left;mso-line-height-rule:exactly;mso-text-raise:2px;">
																			` + message + `</p>
																	</td>
																</tr>
															</table>
														</td>
													</tr>
													<tr>
														<td>
															<table class=t126 role=presentation cellpadding=0 cellspacing=0
																align=center>
																<tr>
																	<!--[if !mso]><!-->
																	<td class=t127 style="width:600px;padding:0 0 34px 0;">
																		<!--<![endif]-->
																		<!--[if mso]><td class=t127 style="width:600px;padding:0 0 34px 0;"><![endif]-->
																		<p class=t133
																			style="margin-bottom:0;Margin-bottom:0;font-family:BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif,'Albert Sans';line-height:22px;font-weight:500;font-style:normal;font-size:14px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#333333;text-align:left;mso-line-height-rule:exactly;mso-text-raise:2px;">
																			Regards,

																		</p>
																		<p class=t133
																		style="margin-bottom:0;Margin-bottom:0;font-family:BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif,'Albert Sans';line-height:22px;font-weight:500;font-style:normal;font-size:14px;text-decoration:none;text-transform:none;letter-spacing:-0.56px;direction:ltr;color:#333333;text-align:left;mso-line-height-rule:exactly;mso-text-raise:2px;">
																		Enyumba Team.

																	</p>
																	</td>

																</tr>

															</table>
														</td>
													</tr>


												</table>
											</td>
										</tr>
									</table>
								</td>
							</tr>
							<tr>
								<td>
									<table class=t29 role=presentation cellpadding=0 cellspacing=0 align=center>
										<tr>
											<!--[if !mso]><!-->
											<td class=t30 style="overflow:hidden;width:540px;padding:40px 30px 40px 30px;">
												<!--<![endif]-->
												<!--[if mso]><td class=t30 style="overflow:hidden;width:600px;padding:40px 30px 40px 30px;"><![endif]-->
												<table role=presentation width=100% cellpadding=0 cellspacing=0>
													<tr>
														<td>
															<table class=t33 role=presentation cellpadding=0 cellspacing=0
																align=center>
																<tr>
																	<!--[if !mso]><!-->
																	<td class=t34 style="width:600px;">
																		<!--<![endif]-->
																		<!--[if mso]><td class=t34 style="width:600px;"><![endif]-->
																		<p class=t40
																			style="margin-bottom:0;Margin-bottom:0;font-family:BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif,'Albert Sans';line-height:22px;font-weight:800;font-style:normal;font-size:18px;text-decoration:none;text-transform:none;letter-spacing:-0.9px;direction:ltr;color:#757575;text-align:center;mso-line-height-rule:exactly;mso-text-raise:1px;">
																			Want updates through more platforms?</p>
																	</td>
																</tr>
															</table>
														</td>
													</tr>
													<tr>
														<td>
															<table class=t43 role=presentation cellpadding=0 cellspacing=0
																align=center>
																<tr>
																	<!--[if !mso]><!-->
																	<td class=t44
																		style="overflow:hidden;width:800px;padding:10px 0 36px 0;">
																		<!--<![endif]-->
																		<!--[if mso]><td class=t44 style="overflow:hidden;width:800px;padding:10px 0 36px 0;"><![endif]-->
																		<div class=t50
																			style="display:inline-table;width:100%;text-align:center;vertical-align:top;">
																			<!--[if mso]>
			<table role=presentation cellpadding=0 cellspacing=0 align=center valign=top><tr><td class=t55 style="width:10px;" width=10></td><td width=24 valign=top><![endif]-->
																			<div class=t56
																				style="display:inline-table;text-align:initial;vertical-align:inherit;width:20%;max-width:44px;">
																				<div class=t57
																					style="padding:0 10px 0 10px;">
																					<table role=presentation width=100%
																						cellpadding=0 cellspacing=0
																						class=t58>
																						<tr>
																							<td class=t59>
																								<div style="font-size:0px;">
																									<img class=t60
																										style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;"
																										width=24 height=24
																										src=https://uploads.tabular.email/e/2feb9749-6369-44a9-90e9-1c26bf36c1a5/90e14628-2d8f-4c64-af7a-410b0a53d60c.png />
																								</div>
																							</td>
																						</tr>
																					</table>
																				</div>
																			</div>
																			<!--[if mso]>
			</td><td class=t55 style="width:10px;" width=10></td><td class=t95 style="width:10px;" width=10></td><td width=24 valign=top><![endif]-->
																			<div class=t96
																				style="display:inline-table;text-align:initial;vertical-align:inherit;width:20%;max-width:44px;">
																				<div class=t97
																					style="padding:0 10px 0 10px;">
																					<table role=presentation width=100%
																						cellpadding=0 cellspacing=0
																						class=t98>
																						<tr>
																							<td class=t99>
																								<div style="font-size:0px;">
																									<img class=t100
																										style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;"
																										width=24 height=24
																										src=https://uploads.tabular.email/e/b158fd0c-1d9a-41bb-885b-099af24afa59/bbde14ea-031f-4dfe-bb34-39af4949882b.png />
																								</div>
																							</td>
																						</tr>
																					</table>
																				</div>
																			</div>
																			<!--[if mso]>
			</td><td class=t95 style="width:10px;" width=10></td><td class=t85 style="width:10px;" width=10></td><td width=24 valign=top><![endif]-->
																			<div class=t86
																				style="display:inline-table;text-align:initial;vertical-align:inherit;width:20%;max-width:44px;">
																				<div class=t87
																					style="padding:0 10px 0 10px;">
																					<table role=presentation width=100%
																						cellpadding=0 cellspacing=0
																						class=t88>
																						<tr>
																							<td class=t89>
																								<div style="font-size:0px;">
																									<img class=t90
																										style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;"
																										width=24 height=24
																										src=https://uploads.tabular.email/e/b158fd0c-1d9a-41bb-885b-099af24afa59/b6f1e7ce-8c7b-41ee-b453-746aaf5e9b57.png />
																								</div>
																							</td>
																						</tr>
																					</table>
																				</div>
																			</div>
																			<!--[if mso]>
			</td><td class=t85 style="width:10px;" width=10></td><td class=t75 style="width:10px;" width=10></td><td width=24 valign=top><![endif]-->
																			<div class=t76
																				style="display:inline-table;text-align:initial;vertical-align:inherit;width:20%;max-width:44px;">
																				<div class=t77
																					style="padding:0 10px 0 10px;">
																					<table role=presentation width=100%
																						cellpadding=0 cellspacing=0
																						class=t78>
																						<tr>
																							<td class=t79>
																								<div style="font-size:0px;">
																									<img class=t80
																										style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;"
																										width=24 height=24
																										src=https://uploads.tabular.email/e/2feb9749-6369-44a9-90e9-1c26bf36c1a5/8cf62035-acff-4f30-bb51-13faa775bd9f.png />
																								</div>
																							</td>
																						</tr>
																					</table>
																				</div>
																			</div>
																			<!--[if mso]>
			</td><td class=t75 style="width:10px;" width=10></td><td class=t65 style="width:10px;" width=10></td><td width=24 valign=top><![endif]-->
																			<div class=t66
																				style="display:inline-table;text-align:initial;vertical-align:inherit;width:20%;max-width:44px;">
																				<div class=t67
																					style="padding:0 10px 0 10px;">
																					<table role=presentation width=100%
																						cellpadding=0 cellspacing=0
																						class=t68>
																						<tr>
																							<td class=t69>
																								<div style="font-size:0px;">
																									<img class=t70
																										style="display:block;border:0;height:auto;width:100%;Margin:0;max-width:100%;"
																										width=24 height=24
																										src=https://uploads.tabular.email/e/b158fd0c-1d9a-41bb-885b-099af24afa59/8e37593e-8033-4bc9-9fee-951849506678.png />
																								</div>
																							</td>
																						</tr>
																					</table>
																				</div>
																			</div>
																			<!--[if mso]>
			</td><td class=t65 style="width:10px;" width=10></td>
			</tr></table>
			<![endif]-->
																		</div>
																	</td>
																</tr>
															</table>
														</td>
													</tr>
													<tr>
														<td>
															<table class=t113 role=presentation cellpadding=0 cellspacing=0
																align=center>
																<tr>
																	<!--[if !mso]><!-->
																	<td class=t114 style="width:600px;">
																		<!--<![endif]-->
																		<!--[if mso]><td class=t114 style="width:600px;"><![endif]-->
																		<p class=t120
																			style="margin-bottom:0;Margin-bottom:0;font-family:BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif,'Albert Sans';line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;direction:ltr;color:#888888;text-align:center;mso-line-height-rule:exactly;mso-text-raise:3px;">
																			Ntinda Complex Bldg Block D, 2rd Floor</p>
																	</td>
																</tr>
															</table>
														</td>
													</tr>
													<tr>
														<td>
															<table class=t103 role=presentation cellpadding=0 cellspacing=0
																align=center>
																<tr>
																	<!--[if !mso]><!-->
																	<td class=t104 style="width:600px;">
																		<!--<![endif]-->
																		<!--[if mso]><td class=t104 style="width:600px;"><![endif]-->
																		<p class=t110
																			style="margin-bottom:0;Margin-bottom:0;font-family:BlinkMacSystemFont,Segoe UI,Helvetica Neue,Arial,sans-serif,'Albert Sans';line-height:22px;font-weight:500;font-style:normal;font-size:12px;text-decoration:none;text-transform:none;direction:ltr;color:#888888;text-align:center;mso-line-height-rule:exactly;mso-text-raise:3px;">
																			<a class=t121 href=https://tabular.email
																				style="margin-bottom:0;Margin-bottom:0;font-weight:700;font-style:normal;text-decoration:none;direction:ltr;color:#888888;mso-line-height-rule:exactly;"
																				target=_blank>Unsubscribe</a> • <a
																				class=t122 href=https://tabular.email
																				style="margin-bottom:0;Margin-bottom:0;font-weight:700;font-style:normal;text-decoration:none;direction:ltr;color:#888888;mso-line-height-rule:exactly;"
																				target=_blank>Privacy policy</a> • <a
																				class=t123 href=https://tabular.email
																				style="margin-bottom:0;Margin-bottom:0;font-weight:700;font-style:normal;text-decoration:none;direction:ltr;color:#878787;mso-line-height-rule:exactly;"
																				target=_blank>Contact us</a>
																		</p>
																	</td>
																</tr>
															</table>
														</td>
													</tr>
												</table>
											</td>
										</tr>
									</table>
								</td>
							</tr>
						</table>
					</td>
				</tr>
			</table>
		</div>
		</body>

		</html>
`,
		Subject: subject,
		ReplyTo: "info@enyumba.com",
	}

	sent, err := client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(sent.Id)

	return nil
}

func GenerateRandomNumber() int64 {
	// Generate 6 bytes of random data
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	// Truncate the random data to 6 digits
	num := int64(b[0])*100000 + int64(b[1])*10000 + int64(b[2])*1000 + int64(b[3])*100 + int64(b[4])*10 + int64(b[5])
	num = num % 900000 // Ensure num is between 0 and 899999
	num += 100000      // Add 100000 to get a number between 100000 and 999999

	return num
}
