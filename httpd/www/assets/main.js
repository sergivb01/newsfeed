$(document).ready(function ($) {
  // TODO: For crazy loading images
  document.querySelectorAll('.hover-me').forEach(el => {
    el.addEventListener('mouseover', () => {
      if (!el.hasAttribute('id')) return

      const id = el.getAttribute('data-target');
      const img = document.getElementById(id);
      const src = img.getAttribute('data-original');

      img.setAttribute("data-original", "")
      img.setAttribute("src", src)
    })
  })
});
