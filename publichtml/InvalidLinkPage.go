package publichtml

// InvalidLinkPage ...
var InvalidLinkPage = `
<!DOCTYPE html>
<!-- This page is displayed when someone navigates to a verify email or reset password link
     but their security token is wrong. This can either mean the user has clicked on a
     stale link (i.e. re-click on a password reset link after resetting their password) or
     (rarely) this could be a sign of a malicious user trying to tamper with your app.
 -->
<html>
  <head>
  <title>Invalid Link</title>
  <style type='text/css'>
    .container {
      border-width: 0px;
      display: block;
      font: inherit;
      font-family: 'Helvetica Neue', Helvetica;
      font-size: 16px;
      height: 30px;
      line-height: 16px;
      margin: 45px 0px 0px 45px;
      padding: 0px 8px 0px 8px;
      position: relative;
      vertical-align: baseline;
    }

    h1, h2, h3, h4, h5 {
      color: #0067AB;
      display: block;
      font: inherit;
      font-family: 'Open Sans', 'Helvetica Neue', Helvetica;
      font-size: 30px;
      font-weight: 600;
      height: 30px;
      line-height: 30px;
      margin: 0 0 15px 0;
      padding: 0 0 0 0;
    }
  </style>
  </head>

  <body> 
    <div class="container">
      <h1>Invalid Link</h1>
    </div> 
  </body>
</html>
`
