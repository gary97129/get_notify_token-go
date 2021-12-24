package main
import "fmt"

const authTmpl = `
<!DOCTYPE html>
<html lang="tw">
    <head>
        <title></title>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <script>
        function oAuth2() {
            var URL = 'https://notify-bot.line.me/oauth/authorize?';
            URL += 'response_type=code';
            URL += '&client_id={{.ClientID}}';
            URL += '&redirect_uri={{.CallbackURL}}';
            URL += '&scope=notify';
            URL += '&state=NO_STATE';
            window.location.href = URL;
        }
    </script>
    </head>
    <body>
        <button onclick="oAuth2();"> 連結到 LineNotify 按鈕 </button>
	</body>
`
func authTmp2(byt []byte) string{
	const html = fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="tw">
	    <head>
		<title>token</title>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
	    </head>
	    <body>
		<h1> %s </h1>
	    </body>
	`,strings(byt))
	return html
}
