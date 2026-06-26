document.querySelectorAll('pre').forEach(pre => {
  const btn = document.createElement('button');
  btn.className = 'copy-code-btn';
  btn.textContent = 'Copy';
  btn.addEventListener('click', () => {
    const code = pre.querySelector('code');
    navigator.clipboard.writeText(code ? code.textContent.trim() : pre.textContent.trim()).then(() => {
      btn.textContent = '✓';
      setTimeout(() => btn.textContent = 'Copy', 2000);
    });
  });
  pre.appendChild(btn);
});
