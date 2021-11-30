resource "aws_lex_bot_alias" "socal_qna_dev" {
  bot_name    = aws_lex_bot.socal_qna.name
  bot_version = aws_lex_bot.socal_qna.version
  description = "Development version of bot to provide SoCal gas customers with help on FAQs"
  name        = "SoCalQnADev"
}
