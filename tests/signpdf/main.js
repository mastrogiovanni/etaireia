import signer from 'node-signpdf';
import fs from 'fs';

const signedPdf = signer.sign(
  fs.readFileSync("/home/michele/Projects/etaireia/approval/static/docs/note.pdf"),
  fs.readFileSync("/home/michele/Projects/etaireia/tests/key.pem"),
);

console.log(signedPdf)
