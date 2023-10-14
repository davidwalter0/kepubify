package kepub

import (
	"testing"
)

func Test_Regexp(t *testing.T) {
	var text = []byte(`
  <p class="p3"><span class="koboSpan" id="kobo.5.1">It was the same words over and over again.</span><span class="koboSpan" id="kobo.5.2">General Pounder, head of Space Exploration Command and Nevada Fields, seemed to find no others.</span><span class="koboSpan" id="kobo.5.3">He stared spellbound at the large screen of the radar object scanner.</span></p>

  <p class="p3"><span class="koboSpan" id="kobo.6.1">After fourteen hours of flight, the Stardust had plunged into the uppermost layers of the earth&#39;s atmosphere.</span><span class="koboSpan" id="kobo.6.2">She is currently starting her third braking ellipse.</span></p>

  <p class="p3"><span class="koboSpan" id="kobo.7.1">The high falling speed had been reduced to a residual value of 5 km/sec while still in empty space.</span><span class="koboSpan" id="kobo.7.2">The performance of the new nuclear chemical engine was by no means overestimated.</span><span class="koboSpan" id="kobo.7.3">The remaining jet mass supply allowed maneuvers that would have been impossible with chemical fuels.</span></p>

  <p class="p3"><span class="koboSpan" id="kobo.8.1">The ship was diverted just above the first air molecules.</span><span class="koboSpan" id="kobo.8.2">The machines worked flawlessly and reliably.</span><span class="koboSpan" id="kobo.8.3">Another breakdown seemed to be impossible.</span></p>


  <p class="p3"><span class="koboSpan" id="kobo.30.1">There was a sentence.He said something.Crest seemed to have grasped what was going on in Rhodan&#39;s mind.</span><span class="koboSpan" id="kobo.30.2">&#34;I&#39;m sorry,&#34; he said weakly.</span><span class="koboSpan" id="kobo.30.3">“It was not at my discretion to circumvent the difficulties.</span><span class="koboSpan" id="kobo.30.4">We were not prepared for your arrival.</span><span class="koboSpan" id="kobo.30.5">According to my information, the third planet of this solar system should be an underdeveloped primordial world with primitive creatures.</span><span class="koboSpan" id="kobo.30.6">You seem to have changed since our last research flight.</span><span class="koboSpan" id="kobo.30.7">“But we didn’t come here to contact you.”</span></p>
`)
	t.Logf("%s\n", AddRegexpSpace(text))
}
