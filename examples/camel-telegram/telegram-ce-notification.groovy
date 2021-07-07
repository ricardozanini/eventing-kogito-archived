from('knative:channel/kogito-channel')
  .convertBodyTo(String.class)
  .to('log:info')
