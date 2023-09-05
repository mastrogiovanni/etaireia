
const { richTextFromMarkdown } = require('@contentful/rich-text-from-markdown');
const { documentToHtmlString } = require('@contentful/rich-text-html-renderer');

async function bootstrap() {
    const document = await richTextFromMarkdown('# Hello World\n## Ciao mamma\n* michele\n* Pluto\n* Topolino');
    const result = await documentToHtmlString(document)
    console.log(result)
}

bootstrap().catch(console.log)
