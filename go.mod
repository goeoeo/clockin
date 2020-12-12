module github.com/phpdi/clockin

go 1.13

require github.com/Comdex/imgo v0.0.0-20200213094239-bb8d436f1e5a

require (
	github.com/phpdi/ant/image v0.0.0
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

replace github.com/phpdi/ant/image => ../ant/image
