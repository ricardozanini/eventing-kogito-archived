import org.apache.camel.builder.RouteBuilder;
import org.apache.camel.component.telegram.TelegramConstants;
import org.apache.camel.component.telegram.TelegramParseMode;

public class TelegramCESandboxRoute extends RouteBuilder {

    @Override
    public void configure() throws Exception {
        from("knative:channel/kogito-channel")
                .convertBodyTo(String.class)
                .setHeader(TelegramConstants.TELEGRAM_PARSE_MODE, constant(TelegramParseMode.MARKDOWN))
                .setHeader(TelegramConstants.TELEGRAM_CHAT_ID, header("ce-chat-id"))
                .to("string-template:/etc/camel/resources/TelegramMessage.tm")
                // uncomment here to use send the message :)
                //.to("telegram:bots?authorizationToken={{authorizationToken}}");
                // comment here to not log anymore
                .to(log("info").showBodyType(false).showBody(true).showExchangePattern(false));
    }
}
